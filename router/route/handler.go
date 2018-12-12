// Package route implements the handler wrapper for each handle.
package route

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) (int, error)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, err := h(w, r)

	if code >= http.StatusBadRequest {
		if err != nil {
			log.Printf("handler error (%d: %s): %s\n",
				code, http.StatusText(code), err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		w.WriteHeader(code)

		jsonErr := json.NewEncoder(w).Encode(err)
		if jsonErr != nil {
			panic(err)
		}
	}
}
