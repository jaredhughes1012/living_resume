package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/rest/mockrest"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/jaredhughes1012/restkit/testrestkit"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	cases := []struct {
		name           string
		ad             *iam.AuthData
		err            error
		expectedStatus int
	}{
		{
			name:           "success",
			ad:             testiam.NewAuthData(),
			err:            nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "conflict",
			ad:             nil,
			err:            iam.ErrAccountExists,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "internal error",
			ad:             nil,
			err:            errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	// Arrange
	input := testiam.NewAccountInput()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockctl := gomock.NewController(t)
			defer mockctl.Finish()
			svc := mockrest.NewMockService(mockctl)

			svc.EXPECT().CreateAccount(gomock.Any(), input).Return(c.ad, c.err)
			w := httptest.NewRecorder()
			r := testrestkit.NewJsonRequest(t, http.MethodPost, "/api/iam/v1/accounts", &input)
			Route(svc).ServeHTTP(w, r)

			assert.Equal(t, c.expectedStatus, w.Code)
			if c.err == nil {
				result := testrestkit.RequireJsonResponse[iam.AuthData](t, w.Result(), c.expectedStatus)
				assert.Equal(t, *c.ad, result)
			}
		})
	}
}

func Test_Authenticate(t *testing.T) {
	cases := []struct {
		name           string
		ad             *iam.AuthData
		err            error
		expectedStatus int
	}{
		{
			name:           "success",
			ad:             testiam.NewAuthData(),
			err:            nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "not found",
			ad:             nil,
			err:            iam.ErrAccountNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "internal error",
			ad:             nil,
			err:            errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	// Arrange
	creds := testiam.NewCredentials()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockctl := gomock.NewController(t)
			defer mockctl.Finish()
			svc := mockrest.NewMockService(mockctl)

			svc.EXPECT().Authenticate(gomock.Any(), creds).Return(c.ad, c.err)
			w := httptest.NewRecorder()
			r := testrestkit.NewJsonRequest(t, http.MethodPost, "/api/iam/v1/authenticate", &creds)
			Route(svc).ServeHTTP(w, r)

			assert.Equal(t, c.expectedStatus, w.Code)
			if c.err == nil {
				result := testrestkit.RequireJsonResponse[iam.AuthData](t, w.Result(), c.expectedStatus)
				assert.Equal(t, *c.ad, result)
			}
		})
	}
}
