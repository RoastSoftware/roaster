// Package user implement the user API endpoint.
package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/gorilla/mux"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

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

	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, "Cannot decode data as user.", http.StatusBadRequest)
		return
	}

	// TODO: Maybe add some kind of helper for empty fields?
	if u.Username == "" || u.Email == "" || u.Password == "" {
		http.Error(w, "Missing fields.", http.StatusBadRequest)
		return
	}

	/* TODO: Enable when database initialization has been added.
	err = model.PutUser(u.User, []byte(u.Password))
	if err != nil {
		// TODO: Implement better error handling.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := u.User
	*/

	// TODO: Answer with a session.
	// log.Println(user)
}

func changeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("PATCH: %s", vars["username"])
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("DELETE: %s", vars["username"])
}

func retrieveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("GET: %s", vars["username"])
}

// Init adds the handlers for the User [/user] endpoint.
func Init(router *mux.Router) {
	// Create User [POST]
	router.HandleFunc("", createUser).Methods(http.MethodPost)

	// View/Handle Specific User [/user/{username}]
	router.HandleFunc("/{username}", changeUser).Methods(http.MethodPatch)
	router.HandleFunc("/{username}", removeUser).Methods(http.MethodDelete)
	router.HandleFunc("/{username}", retrieveUser).Methods(http.MethodGet)
}
