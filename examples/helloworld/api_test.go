package helloworld

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/m/lib/ctest/chttptest"
	"go.coder.com/m/lib/hat"
	"go.coder.com/m/lib/hat/asshat"
)

func TestAPI(tt *testing.T) {
	addr, close := chttptest.StartHTTPServer(tt, &API{})
	defer close()

	t := hat.New(tt, "http://"+addr)

	t.Run("Hello Echo single-parent chain", func(t *hat.T) {
		req := t.Get()

		req.Send(t).Assert(
			t,
			func(t testing.TB, r hat.Response) {
				byt := r.DuplicateBody(t)
				require.Equal(t, "Hello /", string(byt))
			},
		)

		t.Run("underscore", func(t *hat.T) {
			req.Clone(t,
				hat.Path("/1234567890"),
			).Send(t).Assert(
				t,
				func(t testing.TB, r hat.Response) {
					byt := r.DuplicateBody(t)
					require.Equal(t, "Path too long\n", string(byt))
				},
				asshat.StatusEqual(http.StatusBadRequest),
			)
		})
	})
}
