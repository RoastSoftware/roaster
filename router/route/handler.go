// Package route implements the handler wrapper for each handle.
package route

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) (int, error)

type httpWriter struct {
	http.ResponseWriter
	headerWritten bool
}

func (w *httpWriter) WriteHeader(status int) {
	w.headerWritten = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *httpWriter) Write(b []byte) (int, error) {
	w.headerWritten = true
	return w.ResponseWriter.Write(b)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hw := &httpWriter{ResponseWriter: w}

	code, err := h(hw, r)

	if err != nil {
		switch {
		case code >= http.StatusInternalServerError:
			log.Printf(
				"handler error (%d: %s): %+v\n",
				code, http.StatusText(code), err)

		case code >= http.StatusBadRequest:
			log.Printf(
				"handler error (%d: %s): %v\n",
				code, http.StatusText(code), err)
		}

		hw.Header().Set("Content-Type", "application/json; charset=utf-8")
		hw.Header().Set("X-Content-Type-Options", "nosniff")

		// Must be run before json.NewEncoder as writing to the body
		// automatically sets the header to 200 if it's not set yet.
		if !hw.headerWritten {
			hw.WriteHeader(code)
		}

		jsonErr := json.NewEncoder(hw).Encode(err)
		if jsonErr != nil {
			panic(err)
		}

		return
	}

	if !hw.headerWritten {
		hw.WriteHeader(code)
	}
}
