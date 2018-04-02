package hatassert

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// BodyEqual checks if the response body equals expects.
func BodyEqual(t testing.TB, expects []byte) hat.ResponseAssertion {
	return func(r hat.Response) {
		assert.Equal(t, expects, r.DuplicateBody())
	}
}

// BodyMatches ensures the response body matches the regex described by expr.
func BodyMatches(t testing.TB, expr string) hat.ResponseAssertion {
	rg, err := regexp.Compile(expr)
	require.NoError(t, err)

	return func(r hat.Response) {
		byt := r.DuplicateBody()
		if !rg.Match(byt) {
			t.Errorf("body %s does not match expr %v", byt, rg.String())
		}
	}
}
