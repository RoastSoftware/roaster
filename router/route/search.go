// Package route implements the search API endpoint.
package route

import (
	"encoding/json"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func searchAll(w http.ResponseWriter, r *http.Request) (int, error) {
	result, err := model.SearchAll(mux.Vars(r)["query"])
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	if len(result) == 0 {
		return http.StatusNoContent, nil
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Search adds the handlers for the Search [/search] endpoint.
func Search(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// General Search [/search/{query}]
	r.Handle("/{query}", handler(searchAll)).Methods(http.MethodGet)
}
