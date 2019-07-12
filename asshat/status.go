package asshat

import (
	"net/http"
	"testing"

	"go.coder.com/hat"
)

// StatusEqual ensures the response status equals expected.
func StatusEqual(expected int) hat.ResponseAssertion {
	return func(t testing.TB, r hat.Response) {
		t.Helper()

		got := r.StatusCode

		if expected != r.StatusCode {
			t.Errorf("wanted status %v (%v), got status %v (%v)",
				expected, http.StatusText(expected),
				got, http.StatusText(got),
			)
		}
	}
}
