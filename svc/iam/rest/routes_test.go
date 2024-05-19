package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/app"
	"github.com/jaredhughes1012/living_resume/svc/iam/app/mockapp"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/jaredhughes1012/restkit/testrestkit"
	"github.com/stretchr/testify/assert"
)

func runTestRequest(svc app.Service, r *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	Route(svc).ServeHTTP(w, r)
	return w
}

func Test_CreateAccount_Success(t *testing.T) {
	// Arrange
	input := testiam.NewAccountInput()
	ad := testiam.NewAuthData()

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	svc := mockapp.NewMockService(mockctl)

	svc.EXPECT().CreateAccount(gomock.Any(), input).Return(&ad, nil)

	w := runTestRequest(svc, testrestkit.NewJsonRequest(t, http.MethodPost, "/api/iam/v1/accounts", &input))
	result := testrestkit.RequireJsonResponse[iam.AuthData](t, w, http.StatusCreated)
	assert.Equal(t, ad, result)
}

func Test_CreateAccount_Conflict(t *testing.T) {
	input := testiam.NewAccountInput()

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	svc := mockapp.NewMockService(mockctl)

	svc.EXPECT().CreateAccount(gomock.Any(), input).Return(nil, iam.ErrAccountExists)

	w := runTestRequest(svc, testrestkit.NewJsonRequest(t, http.MethodPost, "/api/iam/v1/accounts", &input))
	assert.Equal(t, http.StatusConflict, w.Code)
}
