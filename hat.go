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
// default response assertions.
// Defaults should be explicitly provided to
// Request and Assert.
type T struct {
	*testing.T

	URL    *url.URL
	Client *http.Client
}

// New creates a *T from a *testing.T.
func New(t *testing.T, addr string) *T {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	u, err := url.Parse(addr)
	require.NoError(t, err)

	return &T{
		T:      t,
		URL:    u,
		Client: client,
	}
}

// Run creates a subtest.
// The subtest inherits the settings of T.
func (t *T) Run(name string, fn func(t *T)) {
	t.T.Run(name, func(tt *testing.T) {
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

// Tests represents a group of independent tests.
//
// Nested Tests can be constructed by lifting the
// sub-group with Tests.RunNamed or Tests.RunPaths.
//
//     var myT *T
//
//     Tests{
//         "simple test": func(*T) {},
//
//         "some property holds": Tests{
//             "case 1": func(*T) {},
//             "case 2": func(*T) {},
//         }.RunNamed, // lifts the nested Tests
//     }.RunNamed(myT) // recurse & execute all tests
type Tests map[string]func(*T)

var exampleNestedTest func(*T) = Tests{
	"simple test": func(*T) {},

	"some property holds": Tests{
		"case 1": func(*T) {},
		"case 2": func(*T) {},
	}.RunNamed,

	"some section": Tests{
		"/hello/world": func(*T) {},

		"/foo/bar": Tests{
			"/baz":  func(*T) {},
			"/qux":  func(*T) {},
			"/thud": func(*T) {},
		}.RunPaths,
	}.RunPaths,
}.RunNamed

// RunNamed runs the Tests using (*T).Run.
//
// The order of execution is undefined.
func (ts Tests) RunNamed(t *T) {
	for name, test := range ts {
		t.Run(name, test)
	}
}

// RunPaths runs the Tests using (*T).RunPath.
//
// The order of execution is undefined.
func (ts Tests) RunPaths(t *T) {
	for path, test := range ts {
		t.RunPath(path, test)
	}
}
