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
