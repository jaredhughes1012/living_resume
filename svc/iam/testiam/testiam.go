// Creates fully populated test entities for IAM. Modify as needed for your test
package testiam

import "github.com/jaredhughes1012/living_resume/svc/iam"

// Creates a new test Credentials
func NewCredentials() iam.Credentials {
	return iam.Credentials{
		Email:    "email",
		Password: "password",
	}
}

// Creates a new test AccountInput
func NewAccountInput() iam.AccountInput {
	return iam.AccountInput{
		FirstName:   "firstName",
		LastName:    "lastName",
		Credentials: NewCredentials(),
	}
}

// Creates a new test Identity
func NewIdentity() iam.Identity {
	return iam.Identity{
		AccountId: "accountId",
		Email:     "email",
		FirstName: "firstName",
		LastName:  "lastName",
	}
}

// Creates a new test AuthData
func NewAuthData() iam.AuthData {
	return iam.AuthData{
		Token:    "testToken",
		Identity: NewIdentity(),
	}
}
