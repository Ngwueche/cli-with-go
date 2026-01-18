package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// callbackCatch tries to "catch" a pokemon by comparing a random number
// against a fixed threshold.
func callbackCatch(cfg *config, args ...string) error {
	// Guard clause to ensure we have a name.
	if len(args)  < 1 {
		return errors.New("Please provide a pokemon name.")	
	}
	
	// args[0] is the pokemon name (first argument).
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])

	if err != nil {
		return fmt.Errorf("Error getting location area list: %v", err)
	}

	// Constants are compile-time values; they never change.
	const threshold = 100
	// rand.Intn returns a number in [0, n); here n is BaseExperience.
	randNum := rand.Intn(pokemon.BaseExperience)
	log.Printf("BaseExperience: %d, Random number: %d, Threshold: %d", pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s! It ran away.", pokemon.Name)
	}
	// Map assignment stores the caught pokemon by name.
	cfg.caughtPokemons[pokemon.Name] = pokemon
	fmt.Printf("Successfully caught %s!", pokemon.Name)
	return nil
}
