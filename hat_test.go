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
		dt.URL = "testing"
		assert.NotEqual(t, dt.URL, t.URL)
	})

	t.Run("RunURL Creates deep copy, and appends to URL", func(t T) {
		t.RunPath("/deeper", func(dt T) {
			assert.Equal(t, "http://"+addr+"/deeper", dt.URL)
			assert.NotEqual(t, dt.URL, t.URL)
		})
	})
}
