package asshat

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/m/lib/hat"
)

// HeaderEqual checks that the provided header key and value
// are present in the response.
func HeaderEqual(header, expected string) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		require.Equal(t, expected, r.Header.Get(header))
	}
}
