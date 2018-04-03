package helloworld

import (
	"fmt"
	"net/http"
)

type API struct {
}

func (a *API) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) > 9 {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Body too long")
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello "+req.URL.Path)
}
