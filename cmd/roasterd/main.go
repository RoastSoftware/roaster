// Package main is the starting point for the roasterd web server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/router"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/session"
	"github.com/gorilla/csrf"
)

const (
	portEnvKey            = "PORT"
	databaseSourceEnvKey  = "DATABASE_SOURCE"
	redisAddressEnvKey    = "REDIS_ADDRESS"
	redisPasswordEnvKey   = "REDIS_PASSWORD"
	csrfKeyEnvKey         = "CSRF_KEY"
	sessionKeyHashEnvKey  = "SESSION_HASH_KEY"
	sessionKeyBlockEnvKey = "SESSION_BLOCK_KEY"
)

type flags struct {
	devMode          bool
	address          string
	databaseSource   string
	redisAddress     string
	redisPassword    string
	redisNetwork     string
	sessionHashKey   string
	sessionBlockKey  string
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
	flag.StringVar(&context.sessionHashKey, "session-hash-key", context.sessionHashKey, "Session key used as hash key for secure cookies")
	flag.StringVar(&context.sessionBlockKey, "session-block-key", context.sessionBlockKey, "Session key used as block key for secure cookies")
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

		// Do not require that the CSRF and session keys are set for
		// dev-mode, instead use hardcoded 'insecure' keys.
		context.csrfKey = "insecure-dev-mode-csrf-123456789"
		context.sessionHashKey = "insecure-dev-mode-session-hash01"
		context.sessionBlockKey = "insecure-dev-mode-session-block0"
	}

	if port := os.Getenv(portEnvKey); port != "" {
		context.address = fmt.Sprintf(":%s", port)
	}

	if redisAddress := os.Getenv(redisAddressEnvKey); redisAddress != "" {
		context.redisAddress = redisAddress
	}

	if databaseSource := os.Getenv(databaseSourceEnvKey); databaseSource != "" {
		context.databaseSource = databaseSource
	}

	if redisPassword := os.Getenv(redisPasswordEnvKey); redisPassword != "" {
		context.redisPassword = redisPassword
	}

	if csrfKey := os.Getenv(csrfKeyEnvKey); csrfKey != "" {
		context.csrfKey = csrfKey
	}

	if sessionHashKey := os.Getenv(sessionKeyHashEnvKey); sessionHashKey != "" {
		context.sessionHashKey = sessionHashKey
	}

	if sessionBlockKey := os.Getenv(sessionKeyBlockEnvKey); sessionBlockKey != "" {
		context.sessionBlockKey = sessionBlockKey
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
		!context.devMode, // Disable Secure cookie attribute in dev-mode
		[]byte(context.sessionHashKey), []byte(context.sessionBlockKey))
	if err != nil {
		log.Fatalf("session store returned error: %v", err)
	}

	// Do not require secure verification for CSRF middleware if in
	// dev-mode, such as verifying that the connection goes over HTTPS.
	csrfOpt := csrf.Secure(!context.devMode)

	server := &http.Server{
		Handler:      router.New([]byte(context.csrfKey), csrfOpt),
		Addr:         context.address,
		WriteTimeout: context.writeTimeout * time.Second,
		ReadTimeout:  context.readTimeout * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
