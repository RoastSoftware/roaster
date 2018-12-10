// Package route implements the handler wrapper for each handle.
package route

import (
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) (int, error)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, err := h(w, r)

	if code >= http.StatusBadRequest {
		if err != nil {
			log.Printf("handler error (%d: %s): %v\n",
				code, http.StatusText(code), err)
		}

		http.Error(w, http.StatusText(code), code)
	}
}
