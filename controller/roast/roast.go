// Package user implement the roast API endpoint.
package roast

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

type Snippet struct {
	Language string `json:"type"`
	Code     string `json:"code"`
}

func analyzeCode(w http.ResponseWriter, r *http.Request) {
	var snippet Snippet

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		// TODO: Implement better error handling.
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		http.Error(w, notAuthorizedErr, http.StatusUnauthorized)
	}

	err = json.NewDecoder(r.Body).Decode(&snippet)
	if err != nil {
		log.Println(err)
		http.Error(w, cannotDecodeRoastDataErr, http.StatusBadRequest)
		return
	}

	roast := model.RoastResult{}

	switch snippet.Language {
	case "python3":
		result, err := analyze.WithFlake8(strings.NewReader(snippet.Code))
		if err != nil {
			log.Println(err)
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
		log.Println(result)

		roast.Code = snippet.Code
		roast.Language = snippet.Language
		roast.Username = username
		roast.Score = 100

	default:
		http.Error(w, unsupportedLanguageErr, http.StatusBadRequest)
		return
	}

	err = model.PutRoast(roast)
	if err != nil {
		// TODO: Implement better error handling.
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(roast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
