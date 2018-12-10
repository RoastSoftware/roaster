// Package roast implement the roast API endpoint.
package roast

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
)

const (
	cannotDecodeRoastDataErr = "cannot decode data as roast"
	notAuthorizedErr         = "not authorized"
	internalServerErr        = "internal server error"
	unsupportedLanguageErr   = "unsupported language error"
)

type snippet struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func analyzeCode(w http.ResponseWriter, r *http.Request) {
	var in snippet

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		// TODO: Implement better error handling.
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	// Empty user is handled as an anonymous user (not logged in).
	username, _ := s.Values["username"].(string)

	err = json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		log.Println(err)
		http.Error(w, cannotDecodeRoastDataErr, http.StatusBadRequest)
		return
	}

	var roast *model.RoastResult

	switch in.Language {
	case "python3":
		roast, err = analyze.WithFlake8(username, in.Code)
		if err != nil {
			log.Println(err)
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, unsupportedLanguageErr, http.StatusBadRequest)
		return
	}

	// If the user isn't anonymous, add the result to the database.
	if username != "" {
		err = model.PutRoast(roast)
		if err != nil {
			// TODO: Implement better error handling.
			log.Println(err)
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(roast)
	if err != nil {
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
}

// Init adds the handlers for the Roast [/roast] endpoint.
func Init(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Analyze code snippet [POST].
	r.HandleFunc("", analyzeCode).Methods(http.MethodPost)
}
