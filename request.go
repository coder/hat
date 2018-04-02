package hat

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/stretchr/testify/require"
)

// RequestOption modifies a request.
type RequestOption func(req *http.Request)

// URLParams sets the URL parameters of the request.
func URLParams(v url.Values) RequestOption {
	return func(req *http.Request) {
		req.URL.RawQuery += v.Encode()
	}
}

// Body sets the body of a request.
func Body(r io.Reader) RequestOption {
	rdc, ok := r.(io.ReadCloser)
	if !ok {
		rdc = ioutil.NopCloser(r)
	}

	return func(req *http.Request) {
		req.Body = rdc
	}
}

// Request sends an HTTP request to the endpoint.
// body may be nil.
func (t T) Request(method Method, opts ...RequestOption) *Response {
	req, err := http.NewRequest(string(method), t.URL, nil)
	require.NoError(t, err, "failed to create request")

	for _, opt := range opts {
		opt(req)
	}

	t.Logf("%v %v", req.Method, req.URL)
	resp, err := t.Client.Do(req)
	require.NoError(t, err, "failed to send request")

	return &Response{
		Response: resp,
	}
}

// RequestURL sends a request with endpoint appended to the internal URL.
func (t T) RequestURL(method Method, endpoint string, opts ...RequestOption) *Response {
	opts = append(opts, func(r *http.Request) {
		r.URL.Path += endpoint
	})
	return t.Request(method, opts...)
}
