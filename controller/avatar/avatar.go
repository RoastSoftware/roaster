// Package avatar implements the avatar API endpoint.
package avatar

import (
    "mime/multipart"
    "mime"
	"net/http"
    "io"
    "io/ioutil"
    "strings"
    "log"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
)

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

// Init adds the handlers for the Avatar 
func Init(s *mux.Router) {
    s.Use(middleware.EnforceContentType("multipart/form-data"))
    // Create avatar for user [PUT].
    s.HandleFunc("", createAvatar).Methods(http.MethodPut)

    // Retrieve avatar for user [GET].
    s.HandleFunc("", retrieveAvatar).Methods(http.MethodGet)
}
