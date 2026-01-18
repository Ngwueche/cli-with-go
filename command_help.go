package main

import "fmt"

// callbackHelp prints a list of all registered commands.
func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to WizzyMond")
	fmt.Println("Here are your available commands:")
	availableCommands := getCommands()
	// Range over a map yields key/value pairs; we only need the value (cmd).
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
