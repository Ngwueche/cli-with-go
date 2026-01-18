package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 0 {
			args = cleaned[1:] //indexing from 1 to get the rest of the words as args
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
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
		"exit": {
			name:        "exit",
			description: "Turns off the programme",
			callback:    callbackExit,
		},
	}
}
func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
