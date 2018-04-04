package hatassert

import (
	"regexp"

	"github.com/stretchr/testify/require"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// BodyEqual checks if the response body equals expects.
func BodyEqual(expects []byte) hat.ResponseAssertion {
	return func(t hat.T, r hat.Response) {
		assert.Equal(t, expects, t.DuplicateBody(r))
	}
}

// BodyMatches ensures the response body matches the regex described by expr.
func BodyMatches(expr string) hat.ResponseAssertion {
	// Compiling is expensive, so we do it once.
	rg, err := regexp.Compile(expr)

	return func(t hat.T, r hat.Response) {
		require.NoError(t, err, "failed to compile regex")

		byt := t.DuplicateBody(r)
		if !rg.Match(byt) {
			t.Errorf("body %s does not match expr %v", byt, rg.String())
		}
	}
}
