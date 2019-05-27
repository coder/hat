package hat

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT(tt *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.Copy(rw, req.Body)
	}))
	defer s.Close()

	t := New(tt, s.URL)

	t.Run("Run Creates deep copy", func(dt *T) {
		dt.URL.Path = "testing"
		require.NotEqual(t, dt.URL, t.URL)
	})

	t.Run("RunURL Creates deep copy, and appends to URL", func(t *T) {
		t.RunPath("/deeper", func(dt *T) {
			require.Equal(t, s.URL+"/deeper", dt.URL.String())
			require.NotEqual(t, dt.URL, t.URL)
		})
	})
}
