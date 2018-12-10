// Package route implements the avatar API endpoint.
package route

import (
	"errors"
	"io"
	"io/ioutil"
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

func createAvatar(w http.ResponseWriter, r *http.Request) (int, error) {
	var a model.Avatar

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusUnauthorized, nil
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusUnauthorized, errors.New("not authorized")
	}
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return http.StatusBadRequest, err
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return http.StatusBadRequest, err
			}
			a.Avatar, err = ioutil.ReadAll(p)
			if err != nil {
				return http.StatusBadRequest, err
			}
			a.Username = username
			if len(a.Avatar) >= 10*1000*1000 {
				return http.StatusRequestEntityTooLarge,
					errors.New("too large file (greater than 10MB)")
			}
			err = model.PutAvatar(a)
			if err != nil {
				return http.StatusInternalServerError, err
			}
		}
	} else {
		return http.StatusBadRequest,
			errors.New("received wrong Content-Type, expected multipart/*")
	}

	return http.StatusOK, nil
}

func retrieveAvatar(w http.ResponseWriter, r *http.Request) (int, error) {
	var u string

	vars := mux.Vars(r)
	u = vars["username"]
	avatar, err := model.GetAvatar(u)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "foreign_key_violation" {
			return http.StatusNotFound, nil
		} else {
			return http.StatusInternalServerError, err
		}
	}
	_, err = w.Write(avatar.Avatar)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// Avatar adds the handlers for the Avatar
func Avatar(s *mux.Router) {
	s.Use(middleware.EnforceContentType("multipart/form-data"))

	// Create avatar for user [PUT].
	s.Handle("", handler(createAvatar)).Methods(http.MethodPut)

	// Retrieve avatar for user [GET].
	s.Handle("", handler(retrieveAvatar)).Methods(http.MethodGet)
}
