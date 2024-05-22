package integration

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/rest"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/jaredhughes1012/restkit/testrestkit"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

var client *http.Client
var baseUrl string
var runId string

// Starts up the service and configures a client for testing
func TestMain(m *testing.M) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	ctx := context.Background()
	svc, err := rest.StandardService(ctx)
	if err != nil {
		panic(err)
	} else if err = svc.Setup(ctx, true); err != nil {
		panic(err)
	}

	server := httptest.NewServer(rest.Route(svc))
	client = server.Client()
	defer server.Close()

	baseUrl = server.URL

	runId = os.Getenv("RUN_ID")
	if runId == "" {
		runId = uuid.NewV4().String()
	}

	os.Exit(m.Run())
}
func Test_AccountCreation(t *testing.T) {
	input := testiam.NewAccountInput()
	input.Email = fmt.Sprintf("%s-3@test.com", runId)

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts/initiate?debug=true"), &input)
	code := testrestkit.RequireJsonResponse[iam.ActivationCode](t, res, http.StatusOK)

	idnInput := testiam.NewIdentityInput()
	idnInput.AccountId = runId
	idnInput.Credentials.Email = input.Email
	idnInput.Credentials.Password = "P4ssw0rd!13"
	idnInput.ActivationCode = code.Code
	idnInput.FirstName = fmt.Sprintf("First-%s", runId)
	idnInput.LastName = fmt.Sprintf("Last-%s", runId)

	res = testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts"), &idnInput)
	result := testrestkit.RequireJsonResponse[iam.AuthData](t, res, http.StatusCreated)

	assert.NotEmpty(t, result.Token)
	assert.NotEmpty(t, result.Identity.AccountId)

	// Make sure account cannot be duplicated
	res = testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts/initiate?debug=true"), &input)
	assert.Equal(t, http.StatusConflict, res.StatusCode)

	// Make sure account can be authenticated
	res = testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/authenticate"), &idnInput.Credentials)
	result = testrestkit.RequireJsonResponse[iam.AuthData](t, res, http.StatusOK)

	assert.NotEmpty(t, result.Token)
	assert.NotEmpty(t, result.Identity.AccountId)
}

func Test_AccountCreation_NoCode(t *testing.T) {
	idnInput := testiam.NewIdentityInput()
	idnInput.Credentials.Email = fmt.Sprintf("%s-notFound1@test.com", runId)
	idnInput.Credentials.Password = "P4ssw0rd!13"
	idnInput.ActivationCode = "notFoundCode"

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts"), &idnInput)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_Authenticate_NotFound(t *testing.T) {
	creds := testiam.NewCredentials()
	creds.Email = fmt.Sprintf("%s-notFound2@test.com", runId)
	creds.Password = "P4ssw0rd!13"

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/authenticate"), &creds)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
