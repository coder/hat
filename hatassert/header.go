package hatassert

import (
	"testing"

	"go.coder.com/hat"
	"github.com/stretchr/testify/assert"
)

func HeaderEqual(header, expected string) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		assert.Equal(t, expected, r.Header.Get(header))
	}
}
