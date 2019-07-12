package asshat

import (
	"github.com/stretchr/testify/require"

	"go.coder.com/hat"
)

// HeaderEqual checks that the provided header key and value
// are present in the response.
func HeaderEqual(header, expected string) hat.ResponseAssertion {
	return func(t *hat.T, r hat.Response) {
		t.Helper()
		require.Equal(t, expected, r.Header.Get(header))
	}
}
