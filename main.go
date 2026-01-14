package main

import "github.com/Ngwueche/cli-with-go.git/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	StartRepl(&cfg)

}
