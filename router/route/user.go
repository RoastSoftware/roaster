// Package route implements the user API endpoint.
package route

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
	"github.com/willeponken/causerr"
)

func createUser(w http.ResponseWriter, r *http.Request) (int, error) {
	// Create anonymous struct with model.User and seperated password field.
	// The password will be handled differently and not be a part of the
	// final user.
	u := struct {
		model.User
		Password string `json:"password"`
	}{}
	defer func() {
		u.Password = "" // Empty password field on function return.
	}()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return http.StatusBadRequest, causerr.New(err,
			"Unable to decode request")
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Email == "" || u.Password == "" {
		return http.StatusBadRequest, causerr.New(nil, "Missing field in request")
	}

	err = model.PutUser(u.User, []byte(u.Password))
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	// TODO: Implement auth middleware instead.
	s, _ := session.Get(r, "roaster_auth")
	s.Values["username"] = u.Username

	session.Save(r, w, s)

	err = json.NewEncoder(w).Encode(u.User)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

func changeUser(w http.ResponseWriter, r *http.Request) (int, error) {
	/* TODO
	s, _ := session.Get(r, "roaster_auth")
	if s.Values["username"] == mux.Vars(r)["username"] {

	}
	*/

	return http.StatusNotImplemented, nil
}

func removeUser(w http.ResponseWriter, r *http.Request) (int, error) {
	/* TODO
	s, _ := session.Get(r, "roaster_auth")
	if s.Values["username"] == mux.Vars(r)["username"] {

	}
	*/

	return http.StatusNotImplemented, nil
}

func retrieveUser(w http.ResponseWriter, r *http.Request) (int, error) {
	username := mux.Vars(r)["username"]

	if username == "" {
		return http.StatusBadRequest, causerr.New(
			errors.New("must provide username field"),
			"Missing username parameter in URI")
	}

	user, err := model.GetUser(username)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return http.StatusNotFound, causerr.New(err,
				fmt.Sprintf("User: '%s' doesn't exist", username))
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

func retrieveUserScore(w http.ResponseWriter, r *http.Request) (code int, err error) {
	username := mux.Vars(r)["username"]
	score, err := model.GetUserScore(username)
	if err != nil {
		return http.StatusNotFound, causerr.New(err, "Could not find any score for the provided user")
	}

	err = json.NewEncoder(w).Encode(score)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// User adds the handlers for the User [/user] endpoint.
func User(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Create User [POST].
	r.Handle("", handler(createUser)).Methods(http.MethodPost)

	// View/Handle Specific User [/user/{username}].
	r.Handle("/{username}", handler(changeUser)).Methods(http.MethodPatch)
	r.Handle("/{username}", handler(removeUser)).Methods(http.MethodDelete)
	r.Handle("/{username}", handler(retrieveUser)).Methods(http.MethodGet)

	r.Handle("/{username}/score", handler(retrieveUserScore)).Methods(http.MethodGet)
}
