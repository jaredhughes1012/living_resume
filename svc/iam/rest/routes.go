package rest

import (
	"net/http"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/restkit"
)

func getFilter() *restkit.ErrorFilter {
	return restkit.NewErrorFilter(restkit.ErrorMap{
		iam.ErrAccountExists:   http.StatusConflict,
		iam.ErrAccountNotFound: http.StatusNotFound,
	})
}

// Creates a router with all of the IAM service routes
func Route(svc Service) http.Handler {
	mux := http.NewServeMux()
	errFilter := getFilter()

	mux.HandleFunc("POST /api/iam/v1/accounts", func(w http.ResponseWriter, r *http.Request) {
		var input iam.AccountInput
		if errFilter.WriteIfError(w, restkit.ReadRequestBody(r, &input)) {
			return
		}

		ad, err := svc.CreateAccount(r.Context(), input)
		if !errFilter.WriteIfError(w, err) {
			_ = restkit.WriteResponseJson(w, http.StatusCreated, ad)
		}
	})

	mux.HandleFunc("POST /api/iam/v1/authenticate", func(w http.ResponseWriter, r *http.Request) {
		var creds iam.Credentials
		if errFilter.WriteIfError(w, restkit.ReadRequestBody(r, &creds)) {
			return
		}

		ad, err := svc.Authenticate(r.Context(), creds)
		if !errFilter.WriteIfError(w, err) {
			_ = restkit.WriteResponseJson(w, http.StatusOK, ad)
		}
	})

	return mux
}
