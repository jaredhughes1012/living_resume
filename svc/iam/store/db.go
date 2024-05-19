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

	return m.Up()
}

// MigrateUp implements DB.
func (p postgresDb) MigrateDown(ctx context.Context) error {
	m, err := p.getMigrate()
	if err != nil {
		return err
	}

	return m.Drop()
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

// Transact implements DB.
func (p *postgresDb) Transact(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = fn(ctx, tx); err != nil {
		err = tx.Rollback()
	}

	return err
}

// AddAccount implements DB.
func (p *postgresDb) AddAccount(acct iam.Identity, creds iam.Credentials) {
	createdAt := time.Now()
	var m identity
	m.FromInput(acct, creds, createdAt)

	p.accounts = append(p.accounts, m)
}

// BeginTx implements DB.
func (p *postgresDb) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return p.db.BeginTx(ctx, nil)
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
