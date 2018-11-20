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

	var user model.User

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Cannot decode data as user.", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing fields.", http.StatusBadRequest)
		return
	}

	// TODO
	// model.CreateUser(user)
	log.Println(user)
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
