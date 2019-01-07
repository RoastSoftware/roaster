// Package route implements the feed API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func getBooleanFromString(s string) bool {
	if strings.ToLower(s) == "true" {
		return true
	}

	return false
}

func retrieveFeed(w http.ResponseWriter, r *http.Request) (int, error) {
	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		// Internal server error because the router should have
		// made sure the `page` query is a number.
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	var pageSize uint64 = 25 // Default page size.
	if p := r.URL.Query().Get("page-size"); p != "" {
		pageSize, err = strconv.ParseUint(p, 10, 64)
		if err != nil {
			return http.StatusBadRequest, causerr.New(err,
				"Unable to parse ?page-size as a positive number")
		}
	}

	username := r.URL.Query().Get("user")
	followees := getBooleanFromString(r.URL.Query().Get("followees"))

	if username == "" && followees {
		return http.StatusBadRequest, causerr.New(
			errors.New("request for followees is missing user query parameter"),
			"Friends query parameter also requires the user query parameter")
	}

	feed, err := model.GetFeed(username, followees, page, pageSize)
	if err != nil {
		// Returns Internal Server Error even for non-existing
		// usernames because it'll return an empty feed (and not any
		// error).
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
}
