package hat

import (
	"net/http"
	"net/url"
	"path"
	"testing"
	"time"
)

// T represents a test instance.
// It intentionally does not provide any default request modifiers or
// default response assertions.
// Defaults should be explicitly provided to
// Request and Assert.
type T struct {
	*testing.T

	URL    *url.URL
	Client *http.Client
}

// Make creates a T from a testing.T.
func Make(t *testing.T, addr string) T {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	return T{
		T: t,
		URL: &url.URL{
			Scheme: "http",
			Host:   addr,
		},
		Client: client,
	}
}

// Run creates a subtest.
// The subtest inherits the settings of T.
func (t T) Run(name string, fn func(t T)) {
	t.T.Run(name, func(tt *testing.T) {
		t.T = tt
		u := *t.URL
		t.URL = &u

		fn(t)
	})
}

// RunPath creates a subtest with segment appended to the internal URL.
// It uses segment as the name of the subtest.
func (t T) RunPath(elem string, fn func(t T)) {
	t.Run(elem, func(t T) {
		t.URL.Path = path.Join(t.URL.Path, elem)

		fn(t)
	})
}
