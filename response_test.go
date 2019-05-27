package hat

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/m/lib/ctest/chttptest"
)

func TestBody(t *testing.T) {
	req, err := http.NewRequest("GET", "google.com", nil)
	require.NoError(t, err)

	Body(strings.NewReader("test123"))(t, req)

	byt, err := ioutil.ReadAll(req.Body)
	require.NoError(t, err)

	require.Equal(t, []byte("test123"), byt)
}

func TestResponse(tt *testing.T) {
	tt.Parallel()

	addr, close := chttptest.StartHTTPServer(tt, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.Copy(rw, req.Body)
	}))
	defer close()

	t := New(tt, "http://"+addr)

	req := t.Get(func(t testing.TB, req *http.Request) {
		req.Body = ioutil.NopCloser(strings.NewReader("howdy"))
	})

	t.Run("DuplicateBody", func(t *T) {
		for i := 0; i < 4; i++ {
			require.Equal(t, "howdy", string(req.Clone(t).Send(t).DuplicateBody(t)))
		}
	})

	t.Run("Again", func(t *T) {
		for i := 0; i < 3; i++ {
			t.Logf("Iteration %v", i)
			req.Clone(t,
				func(t testing.TB, req *http.Request) {
					// Ensure request is being copied for every Again.
					req.URL.Path += "/a"
					require.Equal(t, "/a", req.URL.Path)
					req.Body = ioutil.NopCloser(strings.NewReader("a"))
				},
			).Send(t).Assert(
				t,
				func(t testing.TB, resp Response) {
					require.Equal(t, []byte("a"), resp.DuplicateBody(t))
				},
			)
		}
	})
}
