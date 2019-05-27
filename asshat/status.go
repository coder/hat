package asshat

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/m/lib/hat"
)

// StatusEqual ensures the response status equals expected.
func StatusEqual(expected int) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()
		require.Equal(t, expected, r.StatusCode)
	}
}
