package hatassert

import (
	"testing"

	"go.coder.com/hat"
	"github.com/stretchr/testify/assert"
)

// StatusEqual ensures the response status equals expected.
func StatusEqual(expected int) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		assert.Equal(t, expected, r.StatusCode)
	}
}
