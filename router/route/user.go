// Package route implements the user API endpoint.
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
		return http.StatusBadRequest, err
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Email == "" || u.Password == "" {
		return http.StatusBadRequest, errors.New("missing fields")
	}

	err = model.PutUser(u.User, []byte(u.Password))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// TODO: Implement auth middleware instead.
	s, _ := session.Get(r, "roaster_auth")
	s.Values["username"] = u.Username

	session.Save(r, w, s)

	err = json.NewEncoder(w).Encode(u.User)
	if err != nil {
		return http.StatusInternalServerError, err
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
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		return http.StatusBadRequest,
			errors.New("must provide username field")
	}

	user, err := model.GetUser(username)
	if err != nil {
		return http.StatusNotFound, err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, err
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

}
