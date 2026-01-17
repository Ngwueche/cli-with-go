package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args)  < 1 {
		return errors.New("Please provide a pokemon name.")	
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])

	if err != nil {
		return fmt.Errorf("Error getting location area list: %v", err)
	}

	const threshold = 100
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s! It ran away.", pokemon.Name)
	}

	fmt.Printf("Successfully caught %s!", pokemon.Name)
	return nil
}