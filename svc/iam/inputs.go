package iam

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
