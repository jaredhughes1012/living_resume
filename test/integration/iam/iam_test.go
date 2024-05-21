package integration

import (
	"context"
	"fmt"
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
	"github.com/stretchr/testify/require"
)

var client *http.Client
var baseUrl string
var runId string

// Starts up the service and configures a client for testing
func TestMain(m *testing.M) {
	svc, err := rest.StandardService()
	if err != nil {
		panic(err)
	} else if err = svc.Setup(context.Background(), true); err != nil {
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
	input.Credentials.Email = fmt.Sprintf("%s-1@test.com", runId)
	input.Credentials.Password = "P4ssw0rd!13"

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts"), &input)
	result := testrestkit.RequireJsonResponse[iam.AuthData](t, res, http.StatusCreated)
	assert.NotEmpty(t, result.Token)
	assert.NotEmpty(t, result.Identity.AccountId)

	res = testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/authenticate"), &input.Credentials)
	result = testrestkit.RequireJsonResponse[iam.AuthData](t, res, http.StatusOK)
	assert.NotEmpty(t, result.Token)
	assert.NotEmpty(t, result.Identity.AccountId)
}

func Test_AccountCreation_Conflict(t *testing.T) {
	input := testiam.NewAccountInput()
	input.Credentials.Email = fmt.Sprintf("%s-2@test.com", runId)
	input.Credentials.Password = "P4ssw0rd!13"

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts"), &input)
	require.Equal(t, http.StatusCreated, res.StatusCode)

	res = testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/accounts"), &input)
	assert.Equal(t, http.StatusConflict, res.StatusCode)
}

func Test_Authenticate_NotFound(t *testing.T) {
	creds := testiam.NewCredentials()
	creds.Email = fmt.Sprintf("%s-notFound@test.com", runId)
	creds.Password = "P4ssw0rd!13"

	res := testrestkit.DoJsonRequest(t, client, http.MethodPost, testrestkit.ApiUrl(t, baseUrl, "/api/iam/v1/authenticate"), &creds)
	require.Equal(t, http.StatusNotFound, res.StatusCode)
}
