package app

import (
	"context"
	"log/slog"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaredhughes1012/living_resume/svc/iam/authn/mockauthn"
	"github.com/jaredhughes1012/living_resume/svc/iam/store/mockstore"
	"github.com/jaredhughes1012/living_resume/svc/iam/testiam"
	"github.com/stretchr/testify/assert"
)

func Test_Service_CreateAccount(t *testing.T) {
	ctx := context.Background()
	token := "testToken"
	input := testiam.NewAccountInput()

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	db := mockstore.NewMockDB(mockctl)
	issuer := mockauthn.NewMockTokenIssuer(mockctl)

	db.EXPECT().AddAccount(gomock.Any(), input.Credentials)
	issuer.EXPECT().IssueToken(gomock.Any()).Return(token, nil)
	db.EXPECT().Save(ctx).Return(nil)

	svc := NewService(slog.Default(), db, issuer)
	ad, err := svc.CreateAccount(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, token, ad.Token)
	assert.NotEmpty(t, ad.Identity.AccountId)
}

func Test_Service_Authenticate(t *testing.T) {
	ctx := context.Background()
	token := "testToken"
	idn := testiam.NewIdentity()
	creds := testiam.NewCredentials()

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	db := mockstore.NewMockDB(mockctl)
	issuer := mockauthn.NewMockTokenIssuer(mockctl)

	db.EXPECT().FindAccountByCredentials(ctx, creds).Return(&idn, nil)
	issuer.EXPECT().IssueToken(gomock.Any()).Return(token, nil)

	svc := NewService(slog.Default(), db, issuer)
	ad, err := svc.Authenticate(ctx, creds)

	assert.NoError(t, err)
	assert.Equal(t, token, ad.Token)
	assert.NotEmpty(t, ad.Identity.AccountId)
}
