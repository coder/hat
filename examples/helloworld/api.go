package helloworld

import (
	"fmt"
	"net/http"
)

type API struct {
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 9 {
		http.Error(w, "Path too long", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello "+r.URL.Path)
}
