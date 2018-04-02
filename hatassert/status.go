package hatassert

import (
	"testing"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

// StatusEquals ensures the response status equals expected.
func StatusEquals(t testing.TB, expected int) hat.ResponseAssertion {
	return func(r hat.Response) {
		assert.Equal(t, expected, r.StatusCode)
	}
}
