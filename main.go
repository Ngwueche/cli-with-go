package main

import (
	"log"

	"github.com/Ngwueche/cli-with-go.git/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.GetLocationAreaList()
	if err != nil {
		log.Fatalf("Error getting location area list: %v", err)
	}
	log.Printf("Location Areas: %+v", resp)
	//StartRepl()

}
