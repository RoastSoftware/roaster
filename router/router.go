// Package router defines all router/mux handles.
package router

import (
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/router/route"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New returns a router with all handles configured.
func New(csrfKey []byte, csrfOpts ...csrf.Option) http.Handler {
	var h http.Handler

	router := mux.NewRouter()

	// Apply CORS policy.
	router.Use(middleware.CORS())

	// Protect all methods except GET, HEAD, OPTIONS and TRACE with CSRF
	// tokens.
	router.Use(middleware.CSRF(csrfKey, csrfOpts...))

	// User [/user].
	route.User(router.PathPrefix("/user").Subrouter())

	// Roast [/roast].
	route.Roast(router.PathPrefix("/roast").Subrouter())

	// Session [/session].
	route.Session(router.PathPrefix("/session").Subrouter())

	// Feed [/feed].
	route.Feed(router.PathPrefix("/feed").Subrouter())

	// Avatar [/user/{username}/avatar].
	route.Avatar(router.PathPrefix("/user/{username}/avatar").Subrouter())

	// Documentation [/doc].
	route.Static(router.PathPrefix("/doc/").Subrouter(), "/doc/", "./doc")

	// Retrieve the roast.software SPA [GET].
	route.Static(router.PathPrefix("/").Subrouter(), "/", "./www/dist")

	// Redirect HTTP to HTTPS if X-Forward-Proto is set.
	h = middleware.Redirect(router)

	// Recover from panics in handler routines and log the error.
	h = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(h)

	return h
}
