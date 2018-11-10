// Package main is the starting point for the roasterd web server.
package main

import (
	"log"
)

func init() {
	// Log line file:linenumber.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Prefix log output with "[roasterd]".
	log.SetPrefix("[\033[34mroasterd\033[0m] ")
}

func main() {
	log.Println("TODO") // TODO
}
