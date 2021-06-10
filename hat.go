package hat

import (
	"net/http"
	"net/url"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// T represents a test instance.
// It intentionally does not provide any default request modifiers or
// default response requireions.
// Defaults should be explicitly provided to
// Request and Assert.
type T struct {
	*testing.T

	URL    *url.URL
	Client *http.Client
	// persistentOpts are run on every request
	persistentOpts []RequestOption
}

// New creates a *T from a *testing.T.
func New(t *testing.T, baseURL string) *T {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	u, err := url.Parse(baseURL)
	require.NoError(t, err)

	return &T{
		T:      t,
		URL:    u,
		Client: client,
	}
}

func (t *T) AddPersistentOpts(opts ...RequestOption) {
	t.persistentOpts = append(t.persistentOpts, opts...)
}

// Run creates a subtest.
// The subtest inherits the settings of T.
func (t *T) Run(name string, fn func(t *T)) {
	t.T.Run(name, func(tt *testing.T) {
		tt.Helper()
		t := *t
		t.T = tt
		u := *t.URL
		t.URL = &u

		fn(&t)
	})
}

// RunPath creates a subtest with segment appended to the internal URL.
// It uses segment as the name of the subtest.
func (t *T) RunPath(elem string, fn func(t *T)) {
	t.Run(elem, func(t *T) {
		t.URL.Path = path.Join(t.URL.Path, elem)
		// preserve trailing slash
		if elem[len(elem)-1] == '/' && t.URL.Path != "/" {
			t.URL.Path += "/"
		}

		fn(t)
	})
}
