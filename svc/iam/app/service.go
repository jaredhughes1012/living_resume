package app

import (
	"context"
	"log/slog"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/authn"
	"github.com/jaredhughes1012/living_resume/svc/iam/store"
	"github.com/satori/uuid"
)

//go:generate mockgen -destination=./mockapp/service.go -package=mockapp -source=./service.go Service

// Executes logic for the IAM service
type Service interface {
	// Creates a new account
	CreateAccount(ctx context.Context, input iam.AccountInput) (*iam.AuthData, error)

	// Performs all setup required for the service before running
	Setup(ctx context.Context, force bool) error
}

type svc struct {
	db     store.DB
	log    *slog.Logger
	issuer authn.TokenIssuer
}

// Setup implements Service.
func (s *svc) Setup(ctx context.Context, force bool) error {
	if force {
		if err := s.db.MigrateDown(ctx); err != nil {
			return err
		}
	}

	return s.db.MigrateUp(ctx)
}

// CreateAccount implements Service.
func (s *svc) CreateAccount(ctx context.Context, input iam.AccountInput) (*iam.AuthData, error) {
	idn := input.ToIdentity(uuid.NewV4().String())
	log := s.log.With("email", input.Credentials.Email, "accountId", idn.AccountId)

	log.Debug("Adding account to database")
	s.db.AddAccount(idn, input.Credentials)

	log.Debug("Issuing auth token")
	token, err := s.issuer.IssueToken(idn)
	if err != nil {
		return nil, err
	}

	log.Debug("Committing transaction")
	if err := s.db.Save(ctx); err != nil {
		return nil, err
	}

	log.Info("Account created successfully")
	return &iam.AuthData{
		Token:    token,
		Identity: idn,
	}, nil
}

func NewService(log *slog.Logger, db store.DB, issuer authn.TokenIssuer) Service {
	return &svc{
		log:    log,
		db:     db,
		issuer: issuer,
	}
}
