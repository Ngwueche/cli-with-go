package main

import (
	"time"

	"github.com/Ngwueche/cli-with-go.git/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
	caughtPokemons          map[string]pokeapi.Pokemon
}

type exploreConfig struct {
	pokeapiClient pokeapi.Client
	areaName      *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	StartRepl(&cfg)

}
