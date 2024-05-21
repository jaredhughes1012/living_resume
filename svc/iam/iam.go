package iam

import "errors"

var (
	ErrAccountExists   = errors.New("account already exists")
	ErrAccountNotFound = errors.New("account not found")
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
type AccountInput struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Credentials Credentials `json:"credentials"`
}

// Converts this input into an Identity
func (input AccountInput) ToIdentity(acctId string) Identity {
	return Identity{
		AccountId: acctId,
		Email:     input.Credentials.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
}
