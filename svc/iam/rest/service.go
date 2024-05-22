package rest

import (
	"context"
	"log/slog"
	"os"

	"github.com/jaredhughes1012/living_resume/svc/iam"
	"github.com/jaredhughes1012/living_resume/svc/iam/authn"
	"github.com/jaredhughes1012/living_resume/svc/iam/store"
)

//go:generate mockgen -destination=./mockrest/service.go -package=mockrest -source=./service.go Service

// Executes logic for the IAM service
type Service interface {
	// Begin the process of creating an account
	InitiateAccountCreation(ctx context.Context, input iam.AccountInput) (*iam.ActivationCode, error)

	// Creates a new account
	CreateAccount(ctx context.Context, input iam.IdentityInput) (*iam.AuthData, error)

	// Locates an account by credentials and authenticates that identity
	Authenticate(ctx context.Context, creds iam.Credentials) (*iam.AuthData, error)

	// Performs all setup required for the service before running
	Setup(ctx context.Context, force bool) error
}

type svc struct {
	db     store.DB
	log    *slog.Logger
	issuer authn.TokenIssuer
	cache  store.Cache
	codes  authn.CodeGenerator
}

// InitiateAccountCreation implements Service.
func (s *svc) InitiateAccountCreation(ctx context.Context, input iam.AccountInput) (*iam.ActivationCode, error) {
	log := s.log.With("email", input.Email)

	log.Debug("Verifying account does not exist")
	if _, err := s.db.FindAccountByEmail(ctx, input.Email); err == nil {
		return nil, iam.ErrAccountExists
	}

	log.Debug("Generating activation code")
	code, err := s.codes.Generate()
	if err != nil {
		return nil, err
	} else if err = s.cache.CacheActivationCode(ctx, code.Code, input.Email); err != nil {
		return nil, err
	}

	log.Info("Account creation initiated")
	return code, nil
}

// Authenticate implements Service.
func (s *svc) Authenticate(ctx context.Context, creds iam.Credentials) (*iam.AuthData, error) {
	log := s.log.With("email", creds.Email)

	log.Debug("Locating account")
	idn, err := s.db.FindAccountByCredentials(ctx, creds)
	if err != nil {
		return nil, iam.ErrAccountNotFound
	}

	log.Debug("Issuing auth token")
	token, err := s.issuer.IssueToken(*idn)
	if err != nil {
		return nil, err
	}

	log.Info("Account authenticated successfully")
	return &iam.AuthData{
		Token:    token,
		Identity: *idn,
	}, nil
}

// CreateAccount implements Service.
func (s *svc) CreateAccount(ctx context.Context, input iam.IdentityInput) (*iam.AuthData, error) {
	idn := input.ToIdentity()
	log := s.log.With("email", input.Credentials.Email, "accountId", idn.AccountId)

	log.Debug("Verifying activation")
	email, err := s.cache.GetActivationEmail(ctx, input.ActivationCode)
	if err != nil {
		return nil, iam.ErrActivationExpired
	} else if email != input.Credentials.Email {
		return nil, iam.ErrActivationExpired
	}

	log.Debug("Adding account to database")
	s.db.AddAccount(idn, input.Credentials)

	log.Debug("Issuing auth token")
	token, err := s.issuer.IssueToken(idn)
	if err != nil {
		return nil, err
	}

	log.Debug("Committing transaction")
	if err := s.db.Save(ctx); err != nil {
		return nil, iam.ErrAccountExists
	}

	log.Info("Account created successfully")
	return &iam.AuthData{
		Token:    token,
		Identity: idn,
	}, nil
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

func NewService(log *slog.Logger, db store.DB, issuer authn.TokenIssuer, cache store.Cache, codes authn.CodeGenerator) Service {
	return &svc{
		log:    log,
		db:     db,
		issuer: issuer,
		cache:  cache,
		codes:  codes,
	}
}

// Creates a new service using standard configuration
func StandardService(ctx context.Context) (Service, error) {
	db, err := store.NewDB(os.Getenv("POSTGRES_CONNSTR"), os.Getenv("IAM_MIGRATIONS_DIR"))
	if err != nil {
		return nil, err
	}

	cache, err := store.NewCache(ctx, os.Getenv("REDIS_CONNSTR"))
	if err != nil {
		return nil, err
	}

	codes, err := authn.NewCodeGenerator(os.Getenv("CLIENT_URL"))
	if err != nil {
		return nil, err
	}

	issuer := authn.NewJwtTokenIssuer([]byte(os.Getenv("JWT_SECRET")))
	return NewService(slog.Default(), db, issuer, cache, codes), nil
}
