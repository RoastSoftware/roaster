// Package route implements the feed API endpoint.
package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func retrieveGlobal(w http.ResponseWriter, r *http.Request) (int, error) {
	keys, ok := r.URL.Query()["page"]
	if !ok || len(keys) != 1 {
		return http.StatusBadRequest, causerr.New(nil,
			"Must provide '?page=' query parameter with a single number value")
	}

	page, err := strconv.ParseUint(keys[0], 10, 64)
	if err != nil {
		// Internal server error because the router should have
		// made sure the `page` query is a number.
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	feed, err := model.GetGlobalFeed(page)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(feed)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Feed adds the handlers for the Feed [/feed] endpoint.
func Feed(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	r.Handle("", handler(retrieveGlobal)).
		Queries("page", "page:{[0-9]+}").
		Methods(http.MethodGet)

	r.Handle("/{username}", handler(retrieveUser)).
		Queries("page", "page:{[0-9]+}").
		Methods(http.MethodGet)
}
