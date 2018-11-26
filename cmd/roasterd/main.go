// Package main is the starting point for the roasterd web server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/controller"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/csrf"
)

const (
	portEnvKey           = "PORT"
	databaseSourceEnvKey = "DATABASE_SOURCE"
	redisAddressEnvKey   = "REDIS_ADDRESS"
	redisPasswordEnvKey  = "REDIS_PASSWORD"
	csrfKeyEnvKey        = "CSRF_KEY"
	sessionEnvKey        = "SESSION_KEY"
)

type flags struct {
	devMode          bool
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
	flag.StringVar(&context.address, "address", context.address, "Listen address for web server")
	flag.StringVar(&context.databaseSource, "database-source", context.databaseSource, "Database connection source")
	flag.StringVar(&context.redisAddress, "redis-address", context.redisAddress, "Redis instance address")
	flag.StringVar(&context.redisPassword, "redis-password", context.redisPassword, "Redis instance password")
	flag.StringVar(&context.redisNetwork, "redis-network", context.redisNetwork, "Redis instance network type (tcp or udp)")
	flag.UintVar(&context.redisMaxIdleConn, "redis-max-idle-conn", context.redisMaxIdleConn, "Redis max idle connections")
	flag.StringVar(&context.sessionKey, "session-key", context.sessionKey, "Session key used as secret key for secure cookies")
	flag.StringVar(&context.csrfKey, "csrf-key", context.csrfKey, "CSRF key used as secret key for CSRF mitigation")
	flag.BoolVar(&context.devMode, "dev-mode", context.devMode, "Run server in (insecure) development mode")
	flag.Parse()

	// Log line file:linenumber.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Prefix log output with "[roasterd]".
	log.SetPrefix("[\033[34mroasterd\033[0m] ")

	if context.devMode {
		log.Println("WARNING: Running in development mode, using " +
			"insecure CSRF and session keys without verification.")

		// Do not require secure verification for CSRF middleware, such
		// as verifying that the connection goes over HTTPS.
		csrf.Secure(false)

		// Do not require that the CSRF and session keys are set for
		// dev-mode, instead use hardcoded 'insecure' keys.
		context.csrfKey = "insecure-dev-mode-csrf-123456789"
		context.sessionKey = "insecure-dev-mode-session-123456789"
	}

	if port := os.Getenv(portEnvKey); port != "" {
		context.address = fmt.Sprintf(":%s", port)
	}

	if context.databaseSource == "" {
		context.databaseSource = os.Getenv(databaseSourceEnvKey)
	}

	if context.redisAddress == "" {
		context.redisAddress = os.Getenv(redisAddressEnvKey)
	}

	if context.redisPassword == "" {
		context.redisPassword = os.Getenv(redisPasswordEnvKey)
	}

	if context.csrfKey == "" {
		context.csrfKey = os.Getenv(csrfKeyEnvKey)
	}

	if context.sessionKey == "" {
		context.sessionKey = os.Getenv(sessionEnvKey)
	}
}

func main() {

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
