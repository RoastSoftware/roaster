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
)

const (
	portEnvKey           = "PORT"
	databaseSourceEnvKey = "DATABASE_SOURCE"
)

type flags struct {
	address        string
	databaseSource string
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

var context = flags{
	address:      ":5000",
	readTimeout:  15,
	writeTimeout: 15,
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

	// The source argument is prioritized before the environment variable.
	if context.databaseSource == "" {
		if source := os.Getenv(databaseSourceEnvKey); source != "" {
			context.databaseSource = source
		} else {
			log.Fatalln("must provide database source address with credentials")
		}
	}

	err := model.Open(context.databaseSource)
	if err != nil {
		log.Fatalf("database returned error: %v", err)
	}

	controller := controller.New()

	server := &http.Server{
		Handler:      controller,
		Addr:         context.address,
		WriteTimeout: context.writeTimeout * time.Second,
		ReadTimeout:  context.readTimeout * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
