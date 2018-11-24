// Package middleware implements http.Handler middleware.
package middleware

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

var xCSRFToken = http.CanonicalHeaderKey("X-Csrf-Token")

// CSRF wraps the gorilla/csrf Protect middleware with an initiated key and
// injects the X-Csrf-Token.
func CSRF(key []byte) mux.MiddlewareFunc {
	csrf.Secure(false) // TODO Remove this.

	if len(key) < 32 {
		panic("CSRF token must be atleast 32 bytes long")
	}

	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(xCSRFToken, csrf.Token(r))
			h.ServeHTTP(w, r)
		}

		return csrf.Protect(key)(http.HandlerFunc(fn))
	}
}
