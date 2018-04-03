package hat

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.coder.com/ctest/chttptest"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	req, err := http.NewRequest("GET", "google.com", nil)
	require.NoError(t, err)

	Body(strings.NewReader("test123"))(req)

	byt, err := ioutil.ReadAll(req.Body)
	require.NoError(t, err)

	assert.Equal(t, []byte("test123"), byt)
}

func TestResponse(tt *testing.T) {
	tt.Parallel()

	addr, close := chttptest.StartHTTPServer(tt, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.Copy(rw, req.Body)
	}))
	defer close()

	t := New(tt, "http://"+addr)

	resp := t.Request(GET, func(req *http.Request) {
		req.Body = ioutil.NopCloser(strings.NewReader("howdy"))
	})

	t.Run("DuplicateBody", func(t T) {
		for i := 0; i < 4; i++ {
			assert.Equal(t, "howdy", string(DuplicateBody(t, resp)))
		}
	})

	t.Run("But", func(t T) {
		for i := 0; i < 3; i++ {
			t.Logf("Iteration %v", i)
			resp.Again(
				t,
				func(req *http.Request) {
					// Ensure request is being copied for every But.
					req.URL.Path += "/a"
					require.Equal(t, "/a", req.URL.Path)
					req.Body = ioutil.NopCloser(strings.NewReader("a"))
				},
			).Assert(
				func(resp Response) {
					assert.Equal(t, []byte("a"), DuplicateBody(t, resp))
				},
			)
		}
	})
}
