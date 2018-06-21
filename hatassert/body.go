package hatassert

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// BodyEqual checks if the response body equals expects.
// Use BodyStringEqual instead of casting `expects` from a string so
// the error message shows the textual difference.
func BodyEqual(expects []byte) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		assert.Equal(t, expects, r.DuplicateBody(t))
	}
}

// BodyStringEqual checks if the body equals string expects.
func BodyStringEqual(expects string) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		assert.Equal(t, expects, string(r.DuplicateBody(t)))
	}
}

// BodyMatches ensures the response body matches the regex described by expr.
func BodyMatches(expr string) hat.ResponseAssertion {
	// Compiling is expensive, so we do it once.
	rg, err := regexp.Compile(expr)

	return func(t testing.TB, r hat.Response) {
		t.Helper()
		require.NoError(t, err, "failed to compile regex")

		byt := r.DuplicateBody(t)
		if !rg.Match(byt) {
			t.Errorf("body %s does not match expr %v", byt, rg.String())
		}
	}
}
