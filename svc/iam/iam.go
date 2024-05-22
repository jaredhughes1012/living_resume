package iam

import "errors"

var (
	ErrAccountExists     = errors.New("account already exists")
	ErrAccountNotFound   = errors.New("account not found")
	ErrActivationExpired = errors.New("activation code expired")
	ErrMismatchedCode    = errors.New("wrong email provided for activation code")
)

// Represents who a user is and what they can do
type Identity struct {
	AccountId string `json:"accountId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (idn Identity) ToClaims() map[string]any {
	return map[string]any{
		"sub":       idn.AccountId,
		"email":     idn.Email,
		"firstName": idn.FirstName,
		"lastName":  idn.LastName,
	}
}

// Data returned when a user is authenticated
type AuthData struct {
	Token    string   `json:"token"`
	Identity Identity `json:"identity"`
}

// Data used to verify a user's identity
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Creates a new account
type IdentityInput struct {
	ActivationCode string      `json:"activationCode"`
	AccountId      string      `json:"accountId"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Credentials    Credentials `json:"credentials"`
}

// Converts this input into an Identity
func (input IdentityInput) ToIdentity() Identity {
	return Identity{
		AccountId: input.AccountId,
		Email:     input.Credentials.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
}

// Input used to begin creating an account
type AccountInput struct {
	Email string `json:"email"`
}

// Activation code used to verify an account
type ActivationCode struct {
	Code string `json:"code"`
	Url  string `json:"url"`
}
