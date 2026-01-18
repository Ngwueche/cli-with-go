package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Please provide a pokemon name.")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught a pokemon named %s yet.", pokemonName)
	}

	fmt.Printf("pokemon details\n")
	fmt.Printf("Name: %s \n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Specie Name: %s\n", pokemon.Species.Name)
	//fmt.Printf("Abilities: %v", pokemon.Abilities)
	//fmt.Printf("Moves: %v\n", pokemon.Moves)

	fmt.Printf("Stats")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types")
	for _, typ := range pokemon.Types {
		fmt.Printf("Type: - %s", typ.Type.Name)
		fmt.Printf("Slot: - %v \n", typ.Slot)
	}
	return nil
}
