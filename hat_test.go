package hat_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/hat"
	"go.coder.com/hat/asshat"
)

func TestT(tt *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.Copy(rw, req.Body)
	}))
	defer s.Close()

	t := hat.New(tt, s.URL)

	t.Run("Run Creates deep copy", func(dt *hat.T) {
		dt.URL.Path = "testing"
		require.NotEqual(t, dt.URL, t.URL)
	})

	t.Run("RunURL Creates deep copy, and appends to URL", func(t *hat.T) {
		t.RunPath("/deeper", func(dt *hat.T) {
			require.Equal(t, s.URL+"/deeper", dt.URL.String())
			require.NotEqual(t, dt.URL, t.URL)
		})
	})

	t.Run("PersistentOpt", func(t *hat.T) {
		pt := hat.New(tt, s.URL)
		exp := []byte("Hello World!")
		pt.AddPersistentOpts(hat.Body(bytes.NewBuffer(exp)))
		// Opt is attached by persistent opts
		pt.Get().Send(t).Assert(t, asshat.BodyEqual(exp))
	})
}
