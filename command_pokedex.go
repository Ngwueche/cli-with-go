package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	println("Pokemon in PokeDex")
	if len(cfg.caughtPokemons) < 1{
		fmt.Printf("You have not caught any pokemons yet.")
		return nil
	}

	for _, caughtPokemon := range cfg.caughtPokemons {
		fmt.Printf("- %s\n", caughtPokemon.Name)
	}
	
	return nil
}
