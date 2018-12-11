// Package route serves files from the local file system.
package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Static sets up a file server that serves files in a specified directory.
func Static(router *mux.Router, strip, dir string) {
	router.PathPrefix("/").Handler(
		http.StripPrefix(strip,
			http.FileServer(http.Dir(dir))))
}
