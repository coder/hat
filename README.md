# hat

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/go.coder.com/hat)

hat is an HTTP API testing framework for Go.

It's based on composable, reusable response assertions, and request modifiers. It can dramatically **reduce API testing
code**, while **improving clarity of test code and test output**. It leans on the standard `net/http` package
as much as possible.

## Example

Let's test that twitter is working:

```go
func TestTwitter(tt *testing.T) {
    t := hat.New(tt, "https://twitter.com")

    t.Get(
        hat.Path("/realDonaldTrump"),
    ).Send(t).Assert(t,
        asshat.StatusEqual(http.StatusOK),
        asshat.BodyMatches(`President`),
    )
}
```
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [hat](#hat)
  - [Example](#example)
  - [Basic Concepts](#basic-concepts)
    - [Creating Requests](#creating-requests)
    - [Sending Requests](#sending-requests)
    - [Reading Responses](#reading-responses)
  - [Competitive Comparison](#competitive-comparison)
    - [API Symbols](#api-symbols)
    - [LoC](#loc)
    - [net/http](#nethttp)
    - [Chaining APIs](#chaining-apis)
  - [Design Patterns](#design-patterns)
    - [Format Agnostic](#format-agnostic)
    - [Minimal API](#minimal-api)
    - [testing.TB instead of *hat.T](#testingtb-instead-of-hatt)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Basic Concepts

### Creating Requests
hat's entrypoint is its `New` method

```go
func New(t *testing.T, baseURL string) *T
```

which returns a `hat.T` that embeds a `testing.T`, and provides a bunch of methods such as
`Get`, `Post`, and `Patch` to generate HTTP requests. Each request method looks like

```go
func (t *T) Get(opts ...RequestOption) Request
```

RequestOption has the signature

```go
type RequestOption func(t testing.TB, req *http.Request)
```

### Sending Requests

Each request modifies the request however it likes. [A few common `RequestOption`s are provided
in the `hat` package.](https://godoc.org/go.coder.com/hat#RequestOption)

Once the request is built, it can be sent
```go
func (r Request) Send(t *T) *Response
```

or cloned

```go
func (r Request) Clone(t *T, opts ...RequestOption) Request
```
_Cloning is useful when a test is making a slight modification of a complex request._

### Reading Responses

Once you've sent the request, you're given a `hat.Response`. The `Response` should be asserted.

```go
func (r Response) Assert(t testing.TB, assertions ...ResponseAssertion) Response
```

`ResponseAssertion` looks like

```go
type ResponseAssertion func(t testing.TB, r Response)
```

A bunch of pre-made response assertions are available in 
[the `asshat` package](https://godoc.org/go.coder.com/hat/asshat).


## Competitive Comparison

It's difficult to say objectively which framework is the best. But, no existing
framework satisfied us, and we're happy with hat.

| Library                                                    | API Symbols | LoC     | `net/http`               | Custom Assertions/Modifiers |
|------------------------------------------------------------|-------------|---------|--------------------------|-----------------------------|
| hat                                                        | **24**      | **410** | :heavy_check_mark:       | :heavy_check_mark:          |
| [github.com/gavv/httpexpect](//github.com/gavv/httpexpect) | 280         | 10042   | :heavy_multiplication_x: | :warning: (Chaining API)    |
| [github.com/h2non/baloo](//github.com/h2non/baloo)         | 91          | 2146    | :heavy_multiplication_x: | :warning: (Chaining API)    |
| [github.com/h2non/gock](//github.com/h2non/gock)           | 122         | 2957    | :heavy_multiplication_x: | :warning: (Chaining API)    |

_LoC was calculated with cloc._

_Will add more columns and libraries on demand._

### API Symbols

Smaller APIs are easier to use and tend to be less opinionated.

### LoC

Smaller codebases have less bugs and are easier to contribute to.

### net/http

We prefer to use `net/http.Request` and `net/http.Response` so we can reuse the knowledge
we already have. Also, we want to reimplement its surface area.

### Chaining APIs

Chaining APIs look like

```go
 m.GET("/some-path").
        Expect().
        Status(http.StatusOK)
```

We dislike them because they make custom assertions and request modifiers a second-class citizen to
the assertions and modifiers of the package. This encourages the framework's API to bloat,
and discourages abstraction on part of the user.

## Design Patterns

### Format Agnostic

`hat` makes no assumption about the structure of your API, request or response encoding, or
the size of the requests or responses.

### Minimal API

hat and asshat maintains a very small base of helpers. We think of the provided helpers as primitives
for organization and application-specific helpers.

### Always Fatal

While some assertions don't invalidate the test, and would otherwise be `t.Error`s, we don't
really mind when if they fail the test immediately.

To avoid the API complexity of selecting
between `Error`s and `Fatal`s, we fatal all the time.

### testing.TB instead of *hat.T

When porting your code over to hat, it's better to accept a `testing.TB` than a `*hat.T` or a `*testing.T`.

Only accept a `*hat.T` when the function is creating additional requests. This makes the code less coupled,
while clarifying the scope of the helper.

---

This pattern is used in hat itself. The `ResponseAssertion` type and the `Assert` function accept
`testing.TB` instead of a concrete `*hat.T` or `*testing.T`. At first glance, it seems like wherever
the caller is using a `ResponseAssertion` or `Assert`, they would have a `*hat.T`.

In reality, this choice lets consumers hide the initialization of `hat.T` behind a helper function. E.g:

```go
func TestSomething(t *testing.T) {
	makeRequest(t,
		hat.Path("/test"),
	).Assert(t,
		asshat.StatusEqual(t, http.StatusOK),
	)
}
```