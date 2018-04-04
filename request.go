package hat

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/stretchr/testify/require"
)

// RequestOption modifies a request.
type RequestOption func(t T, req *http.Request)

// URLParams sets the URL parameters of the request.
func URLParams(v url.Values) RequestOption {
	return func(_ T, req *http.Request) {
		req.URL.RawQuery += v.Encode()
	}
}

// Path appends path to the URL.
func Path(path string) RequestOption {
	return func(_ T, req *http.Request) {
		req.URL.Path += path
	}
}

// Body sets the body of a request.
func Body(r io.Reader) RequestOption {
	rdc, ok := r.(io.ReadCloser)
	if !ok {
		rdc = ioutil.NopCloser(r)
	}

	return func(_ T, req *http.Request) {
		req.Body = rdc
	}
}

func (t T) sendRequest(req *http.Request) Response {
	t.Logf("%v %v", req.Method, req.URL)
	resp, err := t.Client.Do(req)
	require.NoError(t, err, "failed to send request")

	return Response{
		Response: resp,
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
		copy: copy,
	}
	req.r = req.copy()
	return req
}

// Send dispatches the HTTP request.
func (r Request) Send(t T) Response {
	t.Logf("%v %v", r.r.Method, r.r.URL)
	resp, err := t.Client.Do(r.r)
	require.NoError(t, err, "failed to send request")

	return Response{
		Response: resp,
	}
}

// Clone creates a duplicate HTTP request and applies opts to it.
func (r Request) Clone(t T, opts ...RequestOption) Request {
	return makeRequest(func() *http.Request {
		req := r.copy()
		for _, opt := range opts {
			opt(t, req)
		}
		return req
	})
}

// Request creates an HTTP request to the endpoint.
func (t T) Request(method string, opts ...RequestOption) Request {
	return makeRequest(
		func() *http.Request {
			method = strings.ToUpper(method)

			req, err := http.NewRequest(string(method), t.URL, nil)
			require.NoError(t, err, "failed to create request")

			for _, opt := range opts {
				opt(t, req)
			}

			return req
		},
	)
}

func (t T) Get(opts ...RequestOption) Request {
	return t.Request("Get", opts...)
}

func (t T) Head(opts ...RequestOption) Request {
	return t.Request("Head", opts...)
}

func (t T) Post(opts ...RequestOption) Request {
	return t.Request("Post", opts...)
}

func (t T) Put(opts ...RequestOption) Request {
	return t.Request("Put", opts...)
}

func (t T) Patch(opts ...RequestOption) Request {
	return t.Request("Patch", opts...)
}

func (t T) Delete(opts ...RequestOption) Request {
	return t.Request("Delete", opts...)
}
