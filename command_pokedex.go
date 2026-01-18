package main

import (
	"fmt"
)

// callbackPokedex lists all caught pokemon names.
func callbackPokedex(cfg *config, args ...string) error {
	println("Pokemon in PokeDex")
	// len on a map returns number of key/value pairs.
	if len(cfg.caughtPokemons) < 1{
		fmt.Printf("You have not caught any pokemons yet.")
		return nil
	}

	// Range over a map yields values in random order.
	for _, caughtPokemon := range cfg.caughtPokemons {
		fmt.Printf("- %s\n", caughtPokemon.Name)
	}
	
	return nil
}
