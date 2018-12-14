// Package middleware CORS implementation.
package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS wraps the gorilla/handlers CORS middleware.
func CORS() mux.MiddlewareFunc {
	return handlers.CORS()
}
