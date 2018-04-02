package hat

import (
	"io"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/require"
)

// RequestOption modifies a request.
type RequestOption func(req *http.Request) error

// URLParams sets the URL parameters of the request.
func URLParams(v url.Values) RequestOption {
	return func(req *http.Request) error {
		req.URL.RawQuery += v.Encode()
		return nil
	}
}

// Request sends an HTTP request to the endpoint.
// body may be nil.
func (t T) Request(method Method, body io.Reader, opts ...RequestOption) *Response {
	req, err := http.NewRequest(string(method), t.URL, body)
	require.NoError(t, err, "failed to create request")

	for i, opt := range opts {
		err = opt(req)
		require.NoError(t, err, "failed to apply opt %v", i)
	}

	t.Logf("%v %v", req.Method, req.URL)
	resp, err := t.Client.Do(req)
	require.NoError(t, err, "failed to send request")

	return &Response{
		Response: resp,
	}
}

// RequestURL sends a request with endpoint appended to the internal URL.
func (t T) RequestURL(method Method, endpoint string, body io.Reader, opts ...RequestOption) *Response {
	opts = append(opts, func(r *http.Request) error {
		r.URL.Path += endpoint
		return nil
	})
	return t.Request(method, body, opts...)
}
