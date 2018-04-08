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

	t := hat.Make(tt, addr)

	t.Run("Hello Echo single-parent chain", func(t hat.T) {
		req := t.Get()

		req.Send(t).Assert(
			t,
			func(t testing.TB, r hat.Response) {
				byt := r.DuplicateBody(t)
				assert.Equal(t, "Hello /", string(byt))
			},
		)

		t.Run("underscore", func(t hat.T) {
			req.Clone(
				hat.Path("/1234567890"),
			).Send(t).Assert(
				t,
				func(t testing.TB, r hat.Response) {
					byt := r.DuplicateBody(t)
					assert.Equal(t, "Path too long\n", string(byt))
				},
				hatassert.StatusEqual(http.StatusBadRequest),
			)
		})
	})
}
