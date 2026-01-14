package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreaList(cfg.nextLocationAreaUrl)
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
func callbackMapb(cfg *config) error {
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
