// Package route serves files from the local file system.
package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Static sets up a file server that serves files in a specified directory.
func Static(router *mux.Router, dir string) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
}
