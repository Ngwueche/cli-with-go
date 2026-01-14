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

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the programme",
			callback:    callbackExit,
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
	}
}
func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
