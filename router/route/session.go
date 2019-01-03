// Package route implements the session API endpoint.
package route

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
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
		return http.StatusBadRequest, causerr.New(err, "")
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Password == "" {
		return http.StatusBadRequest, causerr.New(
			errors.New("empty username or password field"),
			"Empty username or password")
	}

	user, ok := model.AuthenticateUser(u.Username, []byte(u.Password))
	if !ok {
		return http.StatusUnauthorized, causerr.New(
			errors.New("invalid username or password"),
			"Invalid username or password")
	}

	// TODO: Implement auth middleware instead.
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}
	s.Values["username"] = user.Username

	session.Save(r, w, s)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

func resumeSession(w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusNoContent, nil
	}

	user, err := model.GetUser(username)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return http.StatusNoContent, nil
		default:
			return http.StatusInternalServerError, causerr.New(err, "")
		}
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

func removeSession(w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	err = session.Invalidate(r, w, s)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

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
	r.Handle("", handler(resumeSession)).Methods(http.MethodGet)

	// Remove Current Session (sign out) [DELETE]
	r.Handle("", handler(removeSession)).Methods(http.MethodDelete)
}
