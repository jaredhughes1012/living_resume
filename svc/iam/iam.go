package iam

import "errors"

var (
	ErrAccountExists   = errors.New("account already exists")
	ErrAccountNotFOund = errors.New("account not found")
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
