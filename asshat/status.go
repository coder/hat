package asshat

import (
	"go.coder.com/hat"
	"net/http"
)

// StatusEqual ensures the response status equals expected.
func StatusEqual(expected int) hat.ResponseAssertion {
	return func(t *hat.T, r hat.Response) {
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
