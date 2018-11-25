// Package session implement the session API endpoint.
package session

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
)

func createSession(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	defer func() {
		u.Password = "" // Empty password field on function return.
	}()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "cannot decode data as user", http.StatusBadRequest)
		return
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Password == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	_, ok := model.AuthenticateUser(u.Username, []byte(u.Password))
	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	// TODO: Implement auth middleware instead.
	s, _ := session.Get(r, "roaster_auth")
	s.Values["username"] = u.Username

	session.Save(r, w, s)
}

func removeSession(w http.ResponseWriter, r *http.Request) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	s.Options.MaxAge = -1 // Set MaxAge to -1 to invalidate session.
	session.Save(r, w, s)
}

// Init adds the handlers for the Session [/session] endpoint.
func Init(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Authenticate for New Session (sign in) [POST]
	r.HandleFunc("", createSession).Methods(http.MethodPost)

	// Remove Current Session (sign out) [DELETE]
	r.HandleFunc("", removeSession).Methods(http.MethodPatch)
}
