// Package redirect checks forward protocol and redirects to https.
package redirect

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	xForwardedProto = http.CanonicalHeaderKey("X-Forwarded-Proto")
)

// Middleware redirects all HTTP requests to HTTPS.
func Middleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get(xForwardedProto)
		if strings.ToUpper(proto) == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf(
				"https://%s%s", r.Host, r.URL,
			), http.StatusPermanentRedirect)

			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
