package authn

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CodeGenerator_Generate(t *testing.T) {
	const testHost = "test.com"

	generator, error := NewCodeGenerator("http://" + testHost)
	require.NoError(t, error)

	code, error := generator.Generate()
	require.NoError(t, error)

	u, err := url.Parse(code.Url)
	require.NoError(t, err)

	assert.Equal(t, codeLength, len(code.Code))
	assert.Equal(t, testHost, u.Host)
	assert.Equal(t, "/api/iam/v1/accounts/activate", u.Path)
	assert.Equal(t, code.Code, u.Query().Get("code"))
}

func Test_CodeGenerator_UniqueCodes(t *testing.T) {
	const testHost = "test.com"

	generator, error := NewCodeGenerator("http://" + testHost)
	require.NoError(t, error)

	code1, error := generator.Generate()
	require.NoError(t, error)
	code2, error := generator.Generate()
	require.NoError(t, error)
	code3, error := generator.Generate()
	require.NoError(t, error)

	assert.NotEqual(t, code1.Code, code2.Code)
	assert.NotEqual(t, code2.Code, code3.Code)
}
