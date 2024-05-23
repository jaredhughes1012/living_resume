package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jaredhughes1012/living_resume/svc/iam"
	_ "github.com/lib/pq"
)

//go:generate mockgen -destination=./mockstore/db.go -package=mockstore -source=./db.go DB

// Manages database interactions
type DB interface {
	// Adds an account to save to the database
	AddAccount(account iam.Identity, creds iam.Credentials)

	// Locates an account by credentials
	FindAccountByCredentials(ctx context.Context, creds iam.Credentials) (*iam.Identity, error)

	// Locates an account by its email address
	FindAccountByEmail(ctx context.Context, email string) (*iam.Identity, error)

	// Rolls back all migrations for this DB
	MigrateDown(ctx context.Context) error

	// Applies all migrations for this DB
	MigrateUp(ctx context.Context) error

	// Saves all changes made to the database
	Save(ctx context.Context) error
}

type postgresDb struct {
	db            *sql.DB
	accounts      []identity
	migrationsDir string
}

// FindAccountByEmail implements DB.
func (p *postgresDb) FindAccountByEmail(ctx context.Context, email string) (*iam.Identity, error) {
	var m identity
	if err := m.findByEmail(ctx, p.db, email); err != nil {
		return nil, err
	}

	idn := m.ToIdentity()
	return &idn, nil
}

// FindAccountByCredentials implements DB.
func (p *postgresDb) FindAccountByCredentials(ctx context.Context, creds iam.Credentials) (*iam.Identity, error) {
	var m identity
	if err := m.findByCredentials(ctx, p.db, creds); err != nil {
		return nil, err
	}

	idn := m.ToIdentity()
	return &idn, nil
}

// Save implements DB.
func (p *postgresDb) Save(ctx context.Context) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, m := range p.accounts {
		if err := m.Save(ctx, tx); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// AddAccount implements DB.
func (p *postgresDb) AddAccount(acct iam.Identity, creds iam.Credentials) {
	createdAt := time.Now()
	var m identity
	m.FromInput(acct, creds, createdAt)

	p.accounts = append(p.accounts, m)
}

func (p postgresDb) getMigrate() (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(p.db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(p.migrationsDir, "postgres", driver)
}

// MigrateUp implements DB.
func (p postgresDb) MigrateUp(ctx context.Context) error {
	m, err := p.getMigrate()
	if err != nil {
		return err
	}

	if err := m.Up(); err == nil || err == migrate.ErrNoChange {
		return nil
	} else {
		return err
	}
}

// MigrateUp implements DB.
func (p postgresDb) MigrateDown(ctx context.Context) error {
	m, err := p.getMigrate()
	if err != nil {
		return err
	}

	if err := m.Down(); err == nil || err == migrate.ErrNoChange {
		return nil
	} else {
		return err
	}
}

// Opens a new database connection
func NewDB(connstr string, migrationsDir string) (DB, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	return &postgresDb{
		db:            db,
		migrationsDir: migrationsDir,
	}, nil
}
