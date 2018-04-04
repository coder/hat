package helloworld

import (
	"fmt"
	"net/http"
)

type API struct {
}

func (a *API) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) > 9 {
		http.Error(rw, "Path too long", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello "+req.URL.Path)
}
