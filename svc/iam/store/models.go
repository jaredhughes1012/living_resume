package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/jaredhughes1012/living_resume/svc/iam"
)

type identity struct {
	Id        string
	CreatedAt time.Time
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (m *identity) FromInput(idn iam.Identity, creds iam.Credentials, createdAt time.Time) {
	m.Id = idn.AccountId
	m.CreatedAt = createdAt
	m.Email = idn.Email
	m.FirstName = idn.FirstName
	m.LastName = idn.LastName
	m.Password = creds.Password
}

func (m identity) ToIdentity() iam.Identity {
	return iam.Identity{
		AccountId: m.Id,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}

func (m identity) Save(ctx context.Context, tx *sql.Tx) error {
	const query = `
    INSERT INTO identity (id, created_at, email, password, first_name, last_name)
    VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT(id) DO UPDATE SET
      email = EXCLUDED.email,
      first_name = EXCLUDED.first_name,
      last_name = EXCLUDED.last_name
  `

	_, err := tx.ExecContext(
		ctx,
		query,
		m.Id,
		m.CreatedAt,
		m.Email,
		m.Password,
		m.FirstName,
		m.LastName,
	)

	return err
}

func (m *identity) findByCredentials(ctx context.Context, db *sql.DB, creds iam.Credentials) error {
	const query = `
    SELECT id, created_at, email, password, first_name, last_name
    FROM identity
    WHERE email = $1
      AND password = $2
  `

	return db.QueryRowContext(ctx, query, creds.Email, creds.Password).Scan(
		&m.Id,
		&m.CreatedAt,
		&m.Email,
		&m.Password,
		&m.FirstName,
		&m.LastName,
	)
}

func (m *identity) findByEmail(ctx context.Context, db *sql.DB, email string) error {
	const query = `
    SELECT id, created_at, email, password, first_name, last_name
    FROM identity
    WHERE email = $1
  `

	return db.QueryRowContext(ctx, query, email).Scan(
		&m.Id,
		&m.CreatedAt,
		&m.Email,
		&m.Password,
		&m.FirstName,
		&m.LastName,
	)
}
