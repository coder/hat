package hat

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// T represents a test instance.
// It intentionally does not provide any default assertions or
// default response modifiers.
// Defaults should be explicitely provided to
// Request and Assert.
type T struct {
	*testing.T
	URL string

	Client *http.Client
}

// New creates a T from a testing.T.
func New(t *testing.T, URL string) *T {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &T{
		T:      t,
		URL:    URL,
		Client: client,
	}
}

// DuplicateBody reads in the response body.
// It replaces the underlying body with a duplicate.
func (t T) DuplicateBody(r Response) []byte {
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	r.Response.Body = ioutil.NopCloser(bytes.NewReader(byt))

	return byt
}

// Run creates a subtest.
// The subtest inherits the settings of T.
func (t T) Run(name string, fn func(t T)) {
	t.T.Run(name, func(tt *testing.T) {
		t.T = tt

		fn(t)
	})
}

// RunPath creates a subtest with segment appended to the internal URL.
// It uses segment as the name of the subtest.
func (t T) RunPath(path string, fn func(t T)) {
	t.Run(path, func(t T) {
		t.URL = urlJoin(t.URL, path)
		fn(t)
	})
}
