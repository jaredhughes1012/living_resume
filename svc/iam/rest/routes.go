package rest

import (
	"net/http"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/app"
	"github.com/jaredhughes1012/restkit"
)

func getFilter() *restkit.ErrorFilter {
	return restkit.NewErrorFilter(restkit.ErrorMap{
		iam.ErrAccountExists: http.StatusConflict,
	})
}

// Creates a router with all of the IAM service routes
func Route(svc app.Service) http.Handler {
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

	return mux
}
