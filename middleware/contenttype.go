// Package middleware implements http.Handler middleware.
package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	contentType = http.CanonicalHeaderKey("Content-Type")
)

// AcceptContentType wraps the Gorilla handlers.ContentTypeHandler as a middleware.
func AcceptContentType(t string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return handlers.ContentTypeHandler(h, contentType, t)
	}
}

// EnforceContentType enforces Content-Type for both the request and response.
func EnforceContentType(t string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		// Set the Content-Type header for outgoing responses.
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(contentType, t)
			h.ServeHTTP(w, r)
		}

		// Ensure that the client sends request with correct
		// Content-Type header set.
		return handlers.ContentTypeHandler(
			http.HandlerFunc(fn),
			contentType, t)
	}
}
