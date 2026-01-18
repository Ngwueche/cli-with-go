package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartRepl runs a basic read-eval-print loop (REPL) that reads stdin,
// parses a command, and dispatches to a callback.
func StartRepl(cfg *config) {
	// NewScanner reads input line-by-line from os.Stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		// Slice indexing to get the command name (first word).
		commandName := cleaned[0]
		// Empty slice literal; will hold remaining words as args.
		args := []string{}
		if len(cleaned) > 0 {
			// Slicing syntax: [1:] means "from index 1 to the end".
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		// Map lookup with "comma ok" to see if the key exists.
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		// Variadic call: args... expands the slice into individual arguments.
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	// Function type field; callback takes *config and a variadic list of args.
	callback    func(cfg *config, args ...string) error
}

// getCommands builds the command registry as a map of name -> metadata.
func getCommands() map[string]cliCommand {
	// Composite literal for map[string]cliCommand with inline struct literals.
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous location areas",
			callback:    callbackMapb,
		},
		"map": {
			name:        "map",
			description: "Displays the location areas",
			callback:    callbackMap,
		},
		"explore": {
			name:        "explore",
			description: "Shows more details about a specific location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch a pokemon",
			description: "catches a pokemon by name",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspects a pokemon",
			description: "inspects a caught pokemon.",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "all caught pokemons",
			description: "lists all caught pokemons.",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the programme",
			callback:    callbackExit,
		},
	}
}
// cleanInput normalizes user input into lowercase tokens.
func cleanInput(str string) []string {

	// strings.ToLower returns a new string; original is unchanged (strings are immutable).
	lowered := strings.ToLower(str)
	// strings.Fields splits on any whitespace and collapses multiple spaces.
	words := strings.Fields(lowered)
	return words

}
