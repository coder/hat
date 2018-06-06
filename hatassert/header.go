package hatassert

import (
	"testing"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
)

func HeaderEqual(header, expected string) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		assert.Equal(t, expected, r.Header.Get(header))
	}
}
