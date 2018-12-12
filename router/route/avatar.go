// Package route implements the avatar API endpoint.
package route

import (
	"errors"
	"fmt"
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
	"github.com/willeponken/causerr"
)

func createAvatar(w http.ResponseWriter, r *http.Request) (int, error) {

	s, err := session.Get(r, "roaster_auth")
	if err != nil {
		return http.StatusUnauthorized, causerr.New(
			nil,
			"Missing or unreadable cookie sent")
	}

	username, ok := s.Values["username"].(string)
	if !ok || username == "" {
		return http.StatusUnauthorized,
			causerr.New(nil, "Unable to authenticate user")
	}
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return http.StatusBadRequest,
			causerr.New(err, "Missing Content-Type header in request")
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return http.StatusBadRequest,
					causerr.New(err, "Unable to read data in request")
			}
			avatar, err := ioutil.ReadAll(p)
			if err != nil {
				return http.StatusBadRequest,
					causerr.New(err, "Unable to read data in request")
			}
			if len(avatar) >= 10*1000*1000 {
				return http.StatusRequestEntityTooLarge,
					causerr.New(
						errors.New("too large file (greater than 10MB)"),
						"Avatar picture must be less or equal to 10 MB")
			}
			a, err := model.NewAvatar(avatar, username)
			if err != nil {
				return http.StatusUnsupportedMediaType,
					causerr.New(err, "Supported image formats are jpeg, png, gif")
			}
			err = model.PutAvatar(a)
			if err != nil {
				return http.StatusInternalServerError,
					causerr.New(err, "")
			}
		}
	} else {
		return http.StatusBadRequest,
			causerr.New(
				errors.New("received wrong Content-Type, expected multipart/*"),
				"Invalid Content-Type, expected multipart/*")
	}

	return http.StatusNoContent, nil
}

func retrieveAvatar(w http.ResponseWriter, r *http.Request) (int, error) {
	var u string

	vars := mux.Vars(r)
	u = vars["username"]
	avatar, err := model.GetAvatar(u)
	if err, ok := err.(*pq.Error); ok {
		if err.Code.Name() == "foreign_key_violation" {
			return http.StatusNotFound,
				causerr.New(nil, fmt.Sprintf("No avatar found for '%s'", u))
		}
		return http.StatusInternalServerError, causerr.New(err, "")
	}
	_, err = w.Write(avatar.Avatar)
	if err != nil {
		return http.StatusInternalServerError, causerr.New(err, "")
	}

	return http.StatusOK, nil
}

// Avatar adds the handlers for the Avatar
func Avatar(s *mux.Router) {
	s.Use(middleware.EnforceContentType("multipart/form-data"))

	// Create avatar for user [PUT].
	s.Handle("", handler(createAvatar)).Methods(http.MethodPut)

	// TODO: implement ability to set GET -
	// content type to image instead of multipart
	// Retrieve avatar for user [GET].
	s.Handle("", handler(retrieveAvatar)).Methods(http.MethodGet)
}
