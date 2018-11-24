// Package main is the starting point for the roasterd web server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/controller"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
)

const (
	portEnvKey           = "PORT"
	databaseSourceEnvKey = "DATABASE_SOURCE"
	csrfTokenEnvKey      = "CSRF_TOKEN"
	sessionEnvKey        = "SESSION_KEY"
)

type flags struct {
	address          string
	databaseSource   string
	redisAddress     string
	redisPassword    string
	redisNetwork     string
	sessionKey       string
	csrfKey          string
	readTimeout      time.Duration
	writeTimeout     time.Duration
	redisMaxIdleConn uint
}

var context = flags{
	address:          ":5000",
	readTimeout:      15,
	writeTimeout:     15,
	redisMaxIdleConn: 10,
	redisAddress:     ":6379",
	redisNetwork:     "tcp",
}

func init() {
	// Log line file:linenumber.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Prefix log output with "[roasterd]".
	log.SetPrefix("[\033[34mroasterd\033[0m] ")
}

func main() {
	if port := os.Getenv(portEnvKey); port != "" {
		context.address = fmt.Sprintf(":%s", port)
	}

	// TODO: Implement flags. The stuff below could be much more DRY.
	if context.databaseSource == "" {
		context.databaseSource = os.Getenv(databaseSourceEnvKey)
	}

	if context.csrfKey == "" {
		context.csrfKey = os.Getenv(csrfTokenEnvKey)
	}

	if context.sessionKey == "" {
		context.sessionKey = os.Getenv(sessionEnvKey)
	}

	err := model.Open(context.databaseSource)
	if err != nil {
		log.Fatalf("database returned error: %v", err)
	}

	err = session.Open(
		context.redisMaxIdleConn,
		context.redisNetwork,
		context.redisAddress,
		context.redisPassword,
		[]byte(context.sessionKey))
	if err != nil {
		log.Fatalf("session store returned error: %v", err)
	}

	controller := controller.New([]byte(context.csrfKey))

	server := &http.Server{
		Handler:      controller,
		Addr:         context.address,
		WriteTimeout: context.writeTimeout * time.Second,
		ReadTimeout:  context.readTimeout * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
