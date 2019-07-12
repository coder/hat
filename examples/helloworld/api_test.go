package helloworld

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go.coder.com/hat"
	"go.coder.com/hat/asshat"
)

func TestAPI(tt *testing.T) {
	s := httptest.NewServer(&API{})
	defer s.Close()

	t := hat.New(tt, s.URL)

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
