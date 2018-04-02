package hat

import (
	"net/http"
	"testing"
)

// T represents a test instance.
// It intentionally does not provide any default assertions or
// default response modifies.
// Defaults should be explicitely provided to
// Request and Assert.
type T struct {
	*testing.T
	URL string

	Client *http.Client
}

// New creates a T from a testing.T.
func New(t *testing.T, URL string) *T {
	return &T{
		T:      t,
		URL:    URL,
		Client: http.DefaultClient,
	}
}

// Run creates a subtest.
// The subtest inherits the settings of T.
func (t T) Run(name string, fn func(t T)) {
	t.T.Run(name, func(tt *testing.T) {
		t.T = tt

		fn(t)
	})
}

// RunURL creates a subtest with segment appended to the internal URL.
// It uses segment as the name of the subtest.
func (t T) RunURL(segment string, fn func(t T)) {
	t.Run(segment, func(t T) {
		t.URL += segment
		fn(t)
	})
}
