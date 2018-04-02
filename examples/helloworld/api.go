package helloworld

import (
	"fmt"
	"net/http"
)

type API struct {
}

func (a *API) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)

	fmt.Fprintf(rw, "Hello "+req.URL.Path)
}
