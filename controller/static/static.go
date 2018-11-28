// Package static serves files from the local file system.
package static

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Init sets up a file server that serves files in a specified directory.
func Init(router *mux.Router, dir string) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
}
