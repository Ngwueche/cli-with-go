package main

import (
	"fmt"
)

// callbackMap fetches and prints the next page of location areas.
func callbackMap(cfg *config, args ...string) error {
	// cfg.nextLocationAreaUrl is a *string; nil means "start from the first page".
	resp, err := cfg.pokeapiClient.GetLocationAreaList(cfg.nextLocationAreaUrl)
	if err != nil {
		return fmt.Errorf("Error getting location area list: %v", err)
	}
	fmt.Printf("Location Areas:")
	// Range over a slice yields index/value; we only need the value.
	for _, locationArea := range resp.Results {
		fmt.Printf("\n - %s", locationArea.Name)
	}
	// Save pagination links for the next/previous commands.
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous
	return nil
}
// callbackMapb fetches and prints the previous page of location areas.
func callbackMapb(cfg *config, args ...string) error {
	// nil check to avoid dereferencing a nil pointer.
	if cfg.previousLocationAreaUrl == nil {
		fmt.Println("No previous page available.")
		return nil
	}
	resp, err := cfg.pokeapiClient.GetLocationAreaList(cfg.previousLocationAreaUrl)
	if err != nil {
		return fmt.Errorf("Error getting location area list: %v", err)
	}
	fmt.Printf("Location Areas:")
	for _, locationArea := range resp.Results {
		fmt.Printf("\n - %s", locationArea.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous
	return nil
}
