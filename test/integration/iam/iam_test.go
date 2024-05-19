package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/rest"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/jaredhughes1012/restkit"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var client *http.Client
var baseUrl *url.URL
var runId string

func apiUrl(t *testing.T, path string) string {
	uPath, err := url.Parse(path)
	require.NoError(t, err)

	return baseUrl.ResolveReference(uPath).String()
}

// Starts up the service and configures a client for testing
func TestMain(m *testing.M) {
	svc, err := rest.NewService()
	if err != nil {
		panic(err)
	} else if err = svc.Setup(context.Background(), true); err != nil {
		panic(err)
	}

	server := httptest.NewServer(rest.Route(svc))
	client = server.Client()
	defer server.Close()

	baseUrl, err = url.Parse(server.URL)
	if err != nil {
		panic(err)
	}

	runId = os.Getenv("RUN_ID")
	if runId == "" {
		runId = uuid.NewV4().String()
	}

	os.Exit(m.Run())
}

func Test_AccountCreation(t *testing.T) {
	input := testiam.NewAccountInput()
	input.Credentials.Email = fmt.Sprintf("%s@test.com", runId)
	input.Credentials.Password = "P4ssw0rd!13"

	req, _ := restkit.NewJsonRequest(http.MethodPost, apiUrl(t, "/api/iam/v1/accounts"), &input)
	res, err := client.Do(req)
	require.NoError(t, err)

	var result iam.AuthData
	require.Equal(t, http.StatusCreated, res.StatusCode)
	require.NoError(t, restkit.ReadJsonResponse(res, &result))

	assert.NotEmpty(t, result.Token)
	assert.NotEmpty(t, result.Identity.AccountId)
}
