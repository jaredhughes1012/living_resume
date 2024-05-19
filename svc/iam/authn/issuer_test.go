package authn

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Issuer_IssueToken(t *testing.T) {
	idn := testiam.NewIdentity()
	hmac := []byte("test")

	issuer := NewJwtTokenIssuer(hmac)
	token, err := issuer.IssueToken(idn)

	require.NoError(t, err)
	tk, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return hmac, nil
	})

	require.NoError(t, err)
	claims, ok := tk.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, idn.AccountId, claims["sub"])
	assert.Equal(t, idn.Email, claims["email"])
	assert.Equal(t, idn.FirstName, claims["firstName"])
	assert.Equal(t, idn.LastName, claims["lastName"])
}
