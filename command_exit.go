package main

import (
	"os"
)

// callbackExit terminates the program immediately.
func callbackExit(cfg *config, args ...string) error {
	// os.Exit skips deferred functions, so this ends the process right away.
	os.Exit(0)
	return nil
}
