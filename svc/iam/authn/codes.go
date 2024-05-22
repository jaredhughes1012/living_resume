package authn

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"

	"github.com/jaredhughes1012/living_resume/svc/iam"
)

const (
	codeChars  = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	codeLength = 6
)

// Generates activation codes for account creation
type CodeGenerator interface {
	Generate() (*iam.ActivationCode, error)
}

type codeGenerator struct {
	activationUrl *url.URL
}

// Generate implements CodeGenerator.
func (c *codeGenerator) Generate() (*iam.ActivationCode, error) {
	bi, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeChars))))
	if err != nil {
		return nil, err
	}

	code := fmt.Sprintf("%0*d", codeLength, bi)
	u, _ := url.Parse(fmt.Sprintf(`/api/iam/v1/accounts/activate?code=%s`, code))

	return &iam.ActivationCode{
		Code: code,
		Url:  c.activationUrl.ResolveReference(u).String(),
	}, nil
}

func NewCodeGenerator(activationUrl string) (CodeGenerator, error) {
	u, err := url.Parse(activationUrl)
	if err != nil {
		return nil, err
	}

	return &codeGenerator{
		activationUrl: u,
	}, nil
}
