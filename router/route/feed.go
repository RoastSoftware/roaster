// Package route implements the feed API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func retrieveFeed(w http.ResponseWriter, r *http.Request) (int, error) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid query parameters"),
			"Must provide '?page=' query parameter with a single number value")
	}

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		// Internal server error because the router should have
		// made sure the `page` query is a number.
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	vars := mux.Vars(r)
	username := vars["username"]

	var feed model.Feed
	if username == "" {
		feed, err = model.GetGlobalFeed(page)
	} else {
		feed, err = model.GetUserFeed(username, page)
	}
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

	// Global Feed [GET].
	r.Handle("", handler(retrieveFeed)).
		Queries("page", "{[0-9]+}").
		Methods(http.MethodGet)

	// User Feed [/feed/{username}].
	r.Handle("/{username}", handler(retrieveFeed)).
		Queries("page", "{[0-9]+}").
		Methods(http.MethodGet)
}
