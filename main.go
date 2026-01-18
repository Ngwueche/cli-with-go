package main

import (
	"time"

	"github.com/Ngwueche/cli-with-go.git/internal/pokeapi"
)

// config holds shared state for the CLI session.
type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
	// map[string]Pokemon lets us look up caught Pokemon by name quickly.
	caughtPokemons          map[string]pokeapi.Pokemon
}

// exploreConfig isn't used yet, but shows how you might bundle params for a command.
type exploreConfig struct {
	pokeapiClient pokeapi.Client
	// *string is a pointer; nil represents "no value".
	areaName      *string
}

func main() {
	// Struct literal with field names gives clarity and is order-independent.
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		// make(map[...]) allocates an empty map ready for use.
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	StartRepl(&cfg)

}
