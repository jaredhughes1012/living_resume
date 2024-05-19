package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
