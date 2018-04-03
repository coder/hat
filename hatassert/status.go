package hatassert

import (
	"testing"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// StatusEqual ensures the response status equals expected.
func StatusEqual(t testing.TB, expected int) hat.ResponseAssertion {
	return func(t hat.T, r hat.Response) {
		assert.Equal(t, expected, r.StatusCode)
	}
}
