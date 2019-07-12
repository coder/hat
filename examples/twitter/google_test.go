package twitter

import (
	"go.coder.com/hat"
	"go.coder.com/hat/asshat"
	"net/http"
	"testing"
)

func TestTwitter(tt *testing.T) {
	t := hat.New(tt, "https://twitter.com")

	t.Get(
		hat.Path("/realDonaldTrump"),
	).Send(t).Assert(t,
		asshat.StatusEqual(http.StatusOK),
		asshat.BodyMatches(`President`),
	)
}
