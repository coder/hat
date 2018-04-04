package hatassert

import (
	"regexp"

	"github.com/stretchr/testify/require"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// BodyEqual checks if the response body equals expects.
func BodyEqual(t hat.T, expects []byte) hat.ResponseAssertion {
	return func(t hat.T, r hat.Response) {
		assert.Equal(t, expects, t.DuplicateBody(r))
	}
}

// BodyMatches ensures the response body matches the regex described by expr.
func BodyMatches(t hat.T, expr string) hat.ResponseAssertion {
	rg, err := regexp.Compile(expr)
	require.NoError(t, err)

	return func(t hat.T, r hat.Response) {
		byt := t.DuplicateBody(r)
		if !rg.Match(byt) {
			t.Errorf("body %s does not match expr %v", byt, rg.String())
		}
	}
}
