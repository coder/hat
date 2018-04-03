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
	rdc, ok := r.(io.ReadCloser)
	if !ok {
		rdc = ioutil.NopCloser(r)
	}

	return func(req *http.Request) {
		req.Body = rdc
	}
}

func (t T) sendRequest(createRequest func() *http.Request) Response {
	req := createRequest()
	if req == nil {
		t.Fatal("nil request created")
	}

	t.Logf("%v %v", req.Method, req.URL)
	resp, err := t.Client.Do(req)
	require.NoError(t, err, "failed to send request")

	return Response{
		Response:      resp,
		createRequest: createRequest,
	}
}

// Request sends an HTTP request to the endpoint.
// body may be nil.
func (t T) Request(method string, opts ...RequestOption) Response {
	method = strings.ToUpper(method)

	return t.sendRequest(func() *http.Request {
		req, err := http.NewRequest(string(method), t.URL, nil)
		require.NoError(t, err, "failed to create request")

		for _, opt := range opts {
			opt(req)
		}
		return req
	})
}

func (t T) Get(opts ...RequestOption) Response {
	return t.Request("Get", opts...)
}

func (t T) Head(opts ...RequestOption) Response {
	return t.Request("Head", opts...)
}

func (t T) Post(opts ...RequestOption) Response {
	return t.Request("Post", opts...)
}

func (t T) Put(opts ...RequestOption) Response {
	return t.Request("Put", opts...)
}

func (t T) Patch(opts ...RequestOption) Response {
	return t.Request("Patch", opts...)
}

func (t T) Delete(opts ...RequestOption) Response {
	return t.Request("Delete", opts...)
}
