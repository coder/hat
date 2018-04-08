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

// Path appends path to the URL.
func Path(path string) RequestOption {
	return func(req *http.Request) {
		req.URL.Path += path
	}
}

// Body sets the body of a request.
func Body(r io.Reader) RequestOption {
	rc, ok := r.(io.ReadCloser)
	if !ok {
		rc = ioutil.NopCloser(r)
	}

	return func(req *http.Request) {
		req.Body = rc
	}
}

// Request represents a pending HTTP request.
type Request struct {
	r *http.Request

	// copy creates an exact copy of the request.
	copy func() *http.Request
}

func makeRequest(copy func() *http.Request) Request {
	req := Request{
		r:    copy(),
		copy: copy,
	}
	return req
}

// Send dispatches the HTTP request.
func (r Request) Send(t T) *Response {
	t.Logf("%v %v", r.r.Method, r.r.URL)

	resp, err := t.Client.Do(r.r)
	require.NoError(t, err, "failed to send request")

	return &Response{
		Response: resp,
	}
}

// Clone creates a duplicate HTTP request and applies opts to it.
func (r Request) Clone(opts ...RequestOption) Request {
	return makeRequest(func() *http.Request {
		req := r.copy()
		for _, opt := range opts {
			opt(req)
		}
		return req
	})
}

// Request creates an HTTP request to the endpoint.
func (t T) Request(method string, opts ...RequestOption) Request {
	return makeRequest(
		func() *http.Request {
			req, err := http.NewRequest(method, t.URL.String(), nil)
			require.NoError(t, err, "failed to create request")

			for _, opt := range opts {
				opt(req)
			}

			return req
		},
	)
}

func (t T) Get(opts ...RequestOption) Request {
	return t.Request(http.MethodGet, opts...)
}

func (t T) Head(opts ...RequestOption) Request {
	return t.Request(http.MethodHead, opts...)
}

func (t T) Post(opts ...RequestOption) Request {
	return t.Request(http.MethodPost, opts...)
}

func (t T) Put(opts ...RequestOption) Request {
	return t.Request(http.MethodPut, opts...)
}

func (t T) Patch(opts ...RequestOption) Request {
	return t.Request(http.MethodPatch, opts...)
}

func (t T) Delete(opts ...RequestOption) Request {
	return t.Request(http.MethodDelete, opts...)
}
