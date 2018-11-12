// Package static serves files from the local file system.
package static

import "net/http"

// NewHandler returns a handler that serves files in a specified directory.
func NewHandler(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}
