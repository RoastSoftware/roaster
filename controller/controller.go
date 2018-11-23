// Package controller defines all router/mux handles.
package controller

import (
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/controller/static"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/controller/user"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New returns a router with all handles configured.
func New() http.Handler {
	var handler http.Handler

	router := mux.NewRouter()

	// User [/user].
	user.Init(router.PathPrefix("/user").Subrouter())

	// Retrieve the roast.software SPA [GET].
	static.Init(router.PathPrefix("/").Subrouter(), "www")

	// Redirect HTTP to HTTPS if X-Forward-Proto is set.
	handler = middleware.Redirect(router)

	// Recover from panics in handler routines and log the error.
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	return handler
}
