package store

import (
	"testing"
	"time"

	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/stretchr/testify/assert"
)

func Test_Account_FromInput(t *testing.T) {
	idn := testiam.NewIdentity()
	creds := testiam.NewCredentials()
	ts := time.Now()

	var m identity
	m.FromInput(idn, creds, ts)

	assert.Equal(t, idn.AccountId, m.Id)
	assert.Equal(t, ts, m.CreatedAt)
	assert.Equal(t, idn.Email, m.Email)
	assert.Equal(t, idn.FirstName, m.FirstName)
	assert.Equal(t, idn.LastName, m.LastName)
	assert.Equal(t, creds.Password, m.Password)
}

func Test_Identity_ToIdentity(t *testing.T) {
	idn := testiam.NewIdentity()
	m := identity{
		Id:        idn.AccountId,
		CreatedAt: time.Now(),
		Email:     idn.Email,
		FirstName: idn.FirstName,
		LastName:  idn.LastName,
	}

	assert.Equal(t, idn, m.ToIdentity())
}
