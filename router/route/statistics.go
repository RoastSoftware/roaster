// Package route implements the route API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func retrieveRoastTimeseries(w http.ResponseWriter, r *http.Request) (int, error) {
	start, err := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid 'start' query parameter"),
			"Missing or invalid 'start' query parameter with timestamp formatted according to RFC3339")
	}

	end, err := time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid 'end' query parameter"),
			"Missing or invalid 'end' query parameter with timestamp formatted according to RFC3339")
	}

	interval, err := time.ParseDuration(fmt.Sprintf("%s",
		r.URL.Query().Get("interval")))
	if err != nil {
		return http.StatusBadRequest, causerr.New(
			errors.New("invalid interval query parameters"),
			"Missing 'interval' query parameter with a time duration formatted such as '300s', '1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'")
	}

	username := r.URL.Query().Get("user")

	var timeseries model.RoastTimeseries
	if username == "" {
		timeseries, err = model.GetGlobalRoastTimeseries(start, end, interval)
	} else {
		timeseries, err = model.GetUserRoastTimeseries(start, end, interval, username)
	}
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(timeseries)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

func retrieveRoastCount(w http.ResponseWriter, r *http.Request) (code int, err error) {
	username := r.URL.Query().Get("user")

	var numberOfRoasts model.NumberOfRoasts
	if username == "" {
		numberOfRoasts, err = model.GetGlobalNumberOfRoasts()
	} else {
		numberOfRoasts, err = model.GetUserNumberOfRoasts(username)
	}
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(numberOfRoasts)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

func retrieveLinesCount(w http.ResponseWriter, r *http.Request) (code int, err error) {
	username := r.URL.Query().Get("user")

	var linesOfCode model.LinesOfCode
	if username == "" {
		linesOfCode, err = model.GetGlobalLinesOfCode()
	} else {
		linesOfCode, err = model.GetUserLinesOfCode(username)
	}
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = json.NewEncoder(w).Encode(linesOfCode)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Statistics adds the handlers for the Statistics [/statistics] endpoint.
func Statistics(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Roast Timeseries [GET].
	r.Handle("/roast/timeseries", handler(retrieveRoastTimeseries)).
		Queries("start", "",
			"end", "",
			"interval", "").
		Methods(http.MethodGet)

	// Roast Count Simple [GET].
	r.Handle("/roast/count", handler(retrieveRoastCount)).
		Methods(http.MethodGet)

	// Lines of Code [GET].
	r.Handle("/roast/lines", handler(retrieveLinesCount)).
		Methods(http.MethodGet)
}
