// Package route implements the session API endpoint.
package route

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
)

func createSession(w http.ResponseWriter, r *http.Request) (int, error) {
	u := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	defer func() {
		u.Password = "" // Empty password field on function return.
	}()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Password == "" {
		return http.StatusBadRequest, errors.New("missing fields")
	}

	user, ok := model.AuthenticateUser(u.Username, []byte(u.Password))
	if !ok {
		return http.StatusUnauthorized, nil
	}

	// TODO: Implement auth middleware instead.
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	s.Values["username"] = u.Username

	session.Save(r, w, s)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func retrieveSession(w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusNoContent, nil
	}

	user, err := model.GetUser(username)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func removeSession(w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	session.Invalidate(r, w, s)

	return http.StatusOK, nil
}

// Session adds the handlers for the Session [/session] endpoint.
func Session(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Authenticate for New Session (sign in) [POST]
	r.Handle("", handler(createSession)).Methods(http.MethodPost)

	// Get Existing Authenticated Session [GET].
	r.Handle("", handler(retrieveSession)).Methods(http.MethodGet)

	// Remove Current Session (sign out) [DELETE]
	r.Handle("", handler(removeSession)).Methods(http.MethodDelete)
}
