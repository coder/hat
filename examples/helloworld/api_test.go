package helloworld

import (
	"net/http"
	"testing"

	"github.com/codercom/hat"
	"github.com/codercom/hat/hatassert"
	"github.com/stretchr/testify/assert"
	"go.coder.com/ctest/chttptest"
)

func TestAPI(tt *testing.T) {
	addr, close := chttptest.StartHTTPServer(tt, &API{})
	defer close()

	t := hat.New(tt, "http://"+addr)

	t.Run("Hello Echo single-parent chain", func(t hat.T) {
		r := t.Get().Assert(
			func(r hat.Response) {
				byt := hat.DuplicateBody(t, r)
				assert.Equal(t, "Hello /", string(byt))
			},
		)

		t.Run("underscore", func(t hat.T) {
			r.Again(t,
				hat.Path("/1234567890"),
			).Assert(
				func(r hat.Response) {
					byt := hat.DuplicateBody(t, r)
					assert.Equal(t, "Path too long", string(byt))
				},
				hatassert.StatusEqual(t, http.StatusBadRequest),
			)
		})
	})
}
