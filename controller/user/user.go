// Package user implement the user API endpoint.
package user

import (
	"encoding/json"
    "mime/multipart"
    "mime"
	"net/http"
    "io"
    "io/ioutil"
    "strings"
    "log"

	// "github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
    "github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
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
        log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Implement auth middleware instead.
	s, _ := session.Get(r, "roaster_auth")
	s.Values["username"] = u.Username

	session.Save(r, w, s)

	err = json.NewEncoder(w).Encode(u.User)
	if err != nil {
        log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func changeUser(w http.ResponseWriter, r *http.Request) {
	/* TODO
	s, _ := session.Get(r, "roaster_auth")
	if s.Values["username"] == mux.Vars(r)["username"] {

	}
	*/

	http.Error(w, "", http.StatusNotImplemented)
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	/* TODO
	s, _ := session.Get(r, "roaster_auth")
	if s.Values["username"] == mux.Vars(r)["username"] {

	}
	*/

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

func createAvatar(w http.ResponseWriter, r *http.Request) {
    var a model.Avatar

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		log.Println(err)
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}
    mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
    if err != nil {
        http.Error(w, "expected a avatar multipart, recieved something else", http.StatusBadRequest)
        return
    }
    if strings.HasPrefix(mediaType, "multipart/") {
        mr := multipart.NewReader(r.Body, params["boundary"])
        for {
            p, err := mr.NextPart()
            if err == io.EOF {
                return
            }
            if err != nil {
                http.Error(w, "error while parsing multipart message", http.StatusBadRequest)
                return
            }
            slurp, err := ioutil.ReadAll(p)
            if err != nil {
                http.Error(w, "error while parsing all in multipart message", http.StatusBadRequest)
            }
            a.Username = username
            a.Avatar = slurp
            err = model.PutAvatar(a)
            if err != nil {
                log.Println(err)
                http.Error(w, "error while inserting image into DB, too large file", http.StatusInternalServerError)
            }
        }
    }
}

func retrieveAvatar(w http.ResponseWriter, r *http.Request) {
    var u string

	vars := mux.Vars(r)
	u = vars["username"]
    avatar, err := model.GetAvatar(u)
    if err != nil {
        http.Error(w, "error when trying to retrieve user avatar from DB", http.StatusNotFound)
        log.Println(err)
        return
    }
    _, err = w.Write(avatar.Avatar)
    if err != nil {
        http.Error(w, "error trying to write avatar to client", http.StatusBadRequest)
        return
    }
}
// Init adds the handlers for the User [/user] endpoint.
func Init(r *mux.Router) {
	// All handlers are required to use application/json as their
	// Content-Type.
    // r.Use(middleware.EnforceContentType("application/json"))

	// Create User [POST].
	r.HandleFunc("", createUser).Methods(http.MethodPost)

	// View/Handle Specific User [/user/{username}].
	r.HandleFunc("/{username}", changeUser).Methods(http.MethodPatch)
	r.HandleFunc("/{username}", removeUser).Methods(http.MethodDelete)
	r.HandleFunc("/{username}", retrieveUser).Methods(http.MethodGet)

    s := r.PathPrefix("/{username}/avatar").Subrouter()
    // s.Use(middleware.EnforceContentType("multipart/form-data"))

    // Create avatar for user [PUT].
    s.HandleFunc("", createAvatar).Methods(http.MethodPut)

    // Retrieve avatar for user [GET].
    s.HandleFunc("", retrieveAvatar).Methods(http.MethodGet)
}
