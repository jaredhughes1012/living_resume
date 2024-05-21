package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Identity_ToClaims(t *testing.T) {
	idn := Identity{
		AccountId: "testId",
		Email:     "test@test.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	result := idn.ToClaims()
	assert.Equal(t, idn.AccountId, result["sub"])
	assert.Equal(t, idn.Email, result["email"])
	assert.Equal(t, idn.FirstName, result["firstName"])
	assert.Equal(t, idn.LastName, result["lastName"])
}

func Test_AccountInput_ToIdentity(t *testing.T) {
	acctId := "testId"
	input := AccountInput{
		FirstName: "John",
		LastName:  "Doe",
		Credentials: Credentials{
			Email:    "test@test.com",
			Password: "password",
		},
	}

	result := input.ToIdentity(acctId)

	assert.Equal(t, acctId, result.AccountId)
	assert.Equal(t, input.Credentials.Email, result.Email)
	assert.Equal(t, input.FirstName, result.FirstName)
	assert.Equal(t, input.LastName, result.LastName)
}
