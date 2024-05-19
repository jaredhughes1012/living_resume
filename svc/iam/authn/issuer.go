package authn

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jaredhughes1012/living_resume/svc/iam"
)

//go:generate mockgen -destination=./mockauthn/issuer.go -package=mockauthn -source=./issuer.go TokenIssuer

type TokenIssuer interface {
	// Issues a token for the given identity
	IssueToken(identity iam.Identity) (string, error)
}

type jwtTokenIssuer struct {
	secret []byte
}

// IssueToken implements TokenIssuer.
func (j *jwtTokenIssuer) IssueToken(identity iam.Identity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(identity.ToClaims()))
	return token.SignedString(j.secret)
}

// NewJwtTokenIssuer creates a new JWT token issuer
func NewJwtTokenIssuer(secret []byte) TokenIssuer {
	return &jwtTokenIssuer{
		secret: secret,
	}
}
