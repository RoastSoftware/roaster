// Package controller defines all router/mux handles.
package controller

import "net/http"
import "github.com/gorilla/mux"
import "github.com/LuleaUniversityOfTechnology/2018-project-roaster/controller/static"

// New returns a router with all handles configured.
func New() *mux.Router {
	mux := mux.NewRouter()
	mux.PathPrefix("/").Handler(http.StripPrefix("", static.NewHandler("www")))

	return mux
}
