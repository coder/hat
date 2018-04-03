package helloworld

import (
	"net/http"
	"testing"

	"github.com/codercom/hat/hatassert"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
	"go.coder.com/ctest/chttptest"
)

func TestAPI(tt *testing.T) {
	addr, close := chttptest.StartHTTPServer(tt, &API{})
	defer close()

	t := hat.New(tt, "http://"+addr)

	t.Run("Hello Echo", func(t hat.T) {
		t.Request(hat.GET).Assert(func(r hat.Response) {
			byt := r.DuplicateBody()
			assert.Equal(t, "Hello /", string(byt))
		}).But(func(req *http.Request) {
			req.URL.Path = "/dog_soup"
		}).Assert(func(r hat.Response) {
			byt := r.DuplicateBody()
			assert.Equal(t, "Hello /dog_soup", string(byt))
		}).But(func(req *http.Request) {
			req.URL.Path = "/1234567890"
		}).Assert(
			func(r hat.Response) {
				byt := r.DuplicateBody()
				assert.Equal(t, "Body too long", string(byt))
			},
			hatassert.StatusEqual(t, http.StatusBadRequest),
		)
	})

	t.Run("Hello Echo single-parent chain", func(t hat.T) {
		t.Request(hat.GET).Assert(
			func(r hat.Response) {
				byt := r.DuplicateBody()
				assert.Equal(t, "Hello /", string(byt))

			},
			func(r hat.Response) {
				r.But(func(req *http.Request) {
					req.URL.Path = "/dog_soup"
				}).Assert(func(r hat.Response) {
					byt := r.DuplicateBody()
					assert.Equal(t, "Hello /dog_soup", string(byt))
				})
			},
		)
	})
}
