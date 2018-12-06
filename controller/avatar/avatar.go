// Package avatar implements the avatar API endpoint.
package avatar

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/middleware"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func createAvatar(w http.ResponseWriter, r *http.Request) {
	var a model.Avatar

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		http.Error(w, "user not logged in", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, "expected an avatar multipart, received something else", http.StatusBadRequest)
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
			a.Avatar, err = ioutil.ReadAll(p)
			if err != nil {
				http.Error(w, "error while parsing all in multipart message", http.StatusBadRequest)
				return
			}
			a.Username = username
			if len(a.Avatar) >= 10*1000*1000 {
				http.Error(w, "too large file (greater than 10MB)", http.StatusRequestEntityTooLarge)
				return
			}
			err = model.PutAvatar(a)
			if err != nil {
				log.Println(err)
				http.Error(w, "error occured when trying to save the image", http.StatusInternalServerError)
				return
			}
		}
	} else {
		http.Error(w, "received wrong Content-Type, expected multipart/*", http.StatusBadRequest)
		return
	}
}

func retrieveAvatar(w http.ResponseWriter, r *http.Request) {
	var u string

	vars := mux.Vars(r)
	u = vars["username"]
	avatar, err := model.GetAvatar(u)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "foreign_key_violation" {
			http.Error(w, "error user doesn't exist", http.StatusNotFound)
			log.Println(err)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	_, err = w.Write(avatar.Avatar)
	if err != nil {
		http.Error(w, "error trying to write avatar to client", http.StatusInternalServerError)
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
