// Package user implement the user API endpoint.
package user

import (
	"encoding/json"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
)

func createUser(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "cannot decode data as user", http.StatusBadRequest)
		return
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Email == "" || u.Password == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	err = model.PutUser(u.User, []byte(u.Password))
	if err != nil {
		// TODO: Implement better error handling.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Password = "" // Empty the password value.
	//user := u.User

	// TODO: Answer with a session.
	// log.Println(user)
}

func changeUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}

func retrieveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "must provide username field", http.StatusBadRequest)
		return
	}

	user, err := model.GetUser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Init adds the handlers for the User [/user] endpoint.
func Init(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
	r.Use(middleware.EnforceContentType("application/json"))

	// Create User [POST].
	r.HandleFunc("", createUser).Methods(http.MethodPost)

	// View/Handle Specific User [/user/{username}].
	r.HandleFunc("/{username}", changeUser).Methods(http.MethodPatch)
	r.HandleFunc("/{username}", removeUser).Methods(http.MethodDelete)
	r.HandleFunc("/{username}", retrieveUser).Methods(http.MethodGet)
}
