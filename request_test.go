package hat_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"go.coder.com/hat/asshat"

	"github.com/stretchr/testify/require"
	"go.coder.com/hat"
)

func TestURLParams(t *testing.T) {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	require.NoError(t, err)

	hat.URLParams(url.Values{
		"q": []string{"sean"},
	})(t, req)

	require.Equal(t, "http://google.com?q=sean", req.URL.String())
}

func TestHeader(tt *testing.T) {
	hKey, hVal := "X-Custom-Header", "test-value"
	// Server returns status code 200 and checks the header is as expected
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		require.Equal(tt, hVal, req.Header.Get(hKey))
		rw.WriteHeader(http.StatusOK)
	}))
	defer s.Close()

	t := hat.New(tt, s.URL)
	t.Get(hat.Header(hKey, hVal)).Send(t).Assert(t, asshat.StatusEqual(http.StatusOK))
}
