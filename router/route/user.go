// Package route implements the user API endpoint.
package route

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
    "log"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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
		return http.StatusBadRequest, causerr.New("missing field in request", "Missing field in request")
	}

	err = model.PutUser(u.User, []byte(u.Password))
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Constraint == "user_email_key" {
				return http.StatusConflict, causerr.New(err, "Email already in use")
			}
			if pgerr.Constraint == "user_pkey" {
				return http.StatusConflict, causerr.New(err, "Username already in use")
			}
		}
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
	vars := mux.Vars(r)
	username := vars["username"]

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

// putFriend returns
func putFriend(w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusNoContent, nil
	}

    f := model.Friend{}

    err = json.NewDecoder(r.Body).Decode(&f)
    if err != nil {
        log.Println(err)
        return http.StatusBadRequest, causerr.New(err,
        "unable to decode request")
    }
    log.Println(username, f)
    if f.Friend == "" {
        return http.StatusBadRequest, causerr.New(
            errors.New("friend is empty"),
            "Missing friend parameter in URI")
    }

    err = model.PutFriend(username, f.Friend)
    if err != nil {
        log.Println(err)
        if pgerr, ok := err.(*pq.Error); ok {
            log.Println(pgerr.Constraint)
            if pgerr.Constraint == "friend_relation_uq" {
                return http.StatusConflict, causerr.New(err, "User already has this friend")
            }
            if pgerr.Constraint == "username_fk" {
                return http.StatusBadRequest, causerr.New(err, "No user registered, are you logged in?")
            }
        }
        return http.StatusInternalServerError, causerr.New(err, "")
    }
    return http.StatusOK, nil
}

// retrieveFriends returns 
func retrieveFriends(w http.ResponseWriter, r *http.Request) (int, error) {
    vars := mux.Vars(r)
    username := vars["username"]

    if username == "" {
        return http.StatusBadRequest, causerr.New(
            errors.New("missing username for lookup"),
            "Missing username for lookup")
    }
    // TODO: make query that either return a friend object full of friends or
    // a list of friend objects, USE LIST OF FRIENDS
    friends, err := model.GetFriends(username)
    if err != nil {
                return http.StatusInternalServerError, causerr.New(err, "")
    }
    if len(friends) <=0 {
        return http.StatusNoContent, causerr.New(err,
        fmt.Sprintf("User: '%s' has no friends", username))

    }

    err = json.NewEncoder(w).Encode(friends)
    if err != nil {
        return http.StatusInternalServerError, causerr.New(err, "error when preparing response")
    }

    return http.StatusOK, nil
}

func removeFriend(w http.ResponseWriter, r *http.Request) (int, error) {
    s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusNoContent, nil
	}

    vars := mux.Vars(r)
    friend := vars["username"]
    // TODO: extract the logged in user in order to know which entry to delete
    if username == "" {
        return http.StatusBadRequest, causerr.New(
            errors.New("missing username for lookup"),
            "Missing username for lookup")
    }
    err = model.RemoveFriend(username, friend)
    if err != nil {
        return http.StatusInternalServerError, causerr.New(err, "error removing friend")
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
    r.Handle("/{username}/friend", handler(putFriend)).Methods(http.MethodPost)
    r.Handle("/{username}/friend", handler(retrieveFriends)).Methods(http.MethodGet)
    r.Handle("/{username}/friend", handler(removeFriend)).Methods(http.MethodDelete)

}
