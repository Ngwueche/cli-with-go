package main

import (
	"errors"
	"fmt"
)

// callbackExplore shows pokemon encounters for a given location area.
func callbackExplore(cfg *config, arg ...string) error {
	// len(arg) checks how many arguments were passed to the variadic parameter.
	if len(arg)  < 1 {
		return errors.New("Please provide an area name to explore")	
	}
	locationAreaName := arg[0]	
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return fmt.Errorf("Error getting location area list: %v", err)
	}
	fmt.Printf("Pokemon in area %s:", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("\n - %s", pokemon.Pokemon.Name)
	}
	
	return nil
}

