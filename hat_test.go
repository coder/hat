package hat

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.coder.com/ctest/chttptest"
)

func TestT(tt *testing.T) {
	addr, close := chttptest.StartHTTPServer(tt, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.Copy(rw, req.Body)
	}))
	defer close()

	t := New(tt, "http://"+addr)

	t.Run("Run Creates deep copy", func(dt T) {
		t.Headers.Add("hello", "world")
		assert.Empty(t, dt.Headers.Get("hello"))
	})

	t.Run("RunURL Creates deep copy, and appends to URL", func(t T) {
		t.RunURL("/deeper", func(dt T) {
			t.Headers.Add("hello2", "world")
			assert.Empty(t, dt.Headers.Get("hello2"))

			assert.Equal(t, "http://"+addr+"/deeper", dt.URL)
		})
	})
}
