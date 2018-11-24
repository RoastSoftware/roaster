// Package util implements helper functions.
package util

import (
	"os"
	"os/signal"
	"syscall"
)

// Graceful calls a function upon program exit.
func Graceful(fn func() error) {
	go func() {
		sig := make(chan os.Signal, 1)
		defer close(sig)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		err := fn()
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}()
}
