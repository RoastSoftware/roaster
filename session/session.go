// Package session implements a session storage using a Redis store and the
// gorilla/sessions package.
package session

import (
	"net/http"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/util"
	"github.com/gorilla/sessions"
	redis "gopkg.in/boj/redistore.v1"
)

var store sessions.Store

// Open a new connection to Redis and initialize a new store.
func Open(maxIdleConn uint, network, address, password string, secure bool, keyPairs ...[]byte) (err error) {
	for _, key := range keyPairs {
		if len(key) != 32 {
			panic("session key(s) for AES-256 must be exactly 32 bytes long")
		}
	}

	rs, err := redis.NewRediStore(int(maxIdleConn), network, address, password, keyPairs...)
	if err != nil {
		return
	}

	defer util.Graceful(rs.Close)

	rs.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
	}

	store = rs

	return
}

// Get returns a session using its name.
func Get(request *http.Request, name string) (*sessions.Session, error) {
	return store.Get(request, name)
}

// New creates a new session by a provided name.
func New(request *http.Request, name string) (*sessions.Session, error) {
	return store.New(request, name)
}

// Save writes a new session.
func Save(request *http.Request, writer http.ResponseWriter, session *sessions.Session) error {
	return store.Save(request, writer, session)
}

// Invalidate sets the max age of a cookie to -1, rendering it invalid.
func Invalidate(request *http.Request, writer http.ResponseWriter, session *sessions.Session) {
	session.Options.MaxAge = -1
	Save(request, writer, session)
}
