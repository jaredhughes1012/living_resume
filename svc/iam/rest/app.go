package rest

import (
	"log/slog"
	"os"

	"github.com/jaredhughes1012/living_resume/svc/iam/app"
	"github.com/jaredhughes1012/living_resume/svc/iam/authn"
	"github.com/jaredhughes1012/living_resume/svc/iam/store"
)

// Creates a new service usable by the rest handler. Uses standard configuration
func NewService() (app.Service, error) {
	db, err := store.NewDB(os.Getenv("POSTGRES_CONNSTR"), os.Getenv("IAM_MIGRATIONS_DIR"))
	if err != nil {
		return nil, err
	}

	issuer := authn.NewJwtTokenIssuer([]byte(os.Getenv("JWT_SECRET")))
	return app.NewService(slog.Default(), db, issuer), nil
}
