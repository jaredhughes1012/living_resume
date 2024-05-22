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
	input := testiam.NewIdentityInput()
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

func Test_InitiateAccountActivation(t *testing.T) {
	cases := []struct {
		name           string
		ac             *iam.ActivationCode
		err            error
		expectedStatus int
		debugQuery     string
	}{
		{
			name:           "no debug",
			ac:             nil,
			err:            nil,
			expectedStatus: http.StatusAccepted,
			debugQuery:     "",
		},
		{
			name:           "debug",
			ac:             testiam.NewActivationCode(),
			err:            nil,
			expectedStatus: http.StatusOK,
			debugQuery:     "true",
		},
		{
			name:           "conflict",
			ac:             nil,
			err:            iam.ErrAccountExists,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "internal error",
			ac:             nil,
			err:            errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "code expired",
			ac:             nil,
			err:            iam.ErrActivationExpired,
			expectedStatus: http.StatusNotFound,
		},
	}

	input := testiam.NewAccountInput()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockctl := gomock.NewController(t)
			defer mockctl.Finish()
			svc := mockrest.NewMockService(mockctl)

			svc.EXPECT().InitiateAccountCreation(gomock.Any(), *input).Return(tc.ac, tc.err)
			w := httptest.NewRecorder()
			r := testrestkit.NewJsonRequest(t, http.MethodPost, "/api/iam/v1/accounts/initiate?debug="+tc.debugQuery, &input)
			Route(svc).ServeHTTP(w, r)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.ac != nil {
				result := testrestkit.RequireJsonResponse[iam.ActivationCode](t, w.Result(), tc.expectedStatus)
				assert.Equal(t, *tc.ac, result)
			}
		})
	}
}
