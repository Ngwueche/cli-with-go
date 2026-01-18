package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemonList fetches a paginated list of pokemon.
func (c *Client) GetPokemonList(pageUrl *string) (PokemonList, error) {
	endpoint := "pokemon/"
	fullUrl := pokeBaseURL + endpoint

	// If a page URL is provided, use it instead of the default endpoint.
	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	// Check the cache before doing a network call.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit for", fullUrl)
		pokemonListResponse := PokemonList{}
		err := json.Unmarshal(dat, &pokemonListResponse)
		if err != nil {
			return PokemonList{}, err
		}
		return pokemonListResponse, nil
	}
	fmt.Println("Cache miss for", fullUrl)
	req, error := http.NewRequest("GET", fullUrl, nil)
	if error != nil {
		return PokemonList{}, error
	}

	res, error := c.httpClient.Do(req)
	if error != nil {
		return PokemonList{}, error
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokemonList{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonList{}, err
	}

	pokemonListResponse := PokemonList{}
	err = json.Unmarshal(data, &pokemonListResponse)
	if err != nil {
		return PokemonList{}, err
	}
	c.cache.Add(fullUrl, data)
	return pokemonListResponse, nil
}

// GetPokemon fetches full details for a single pokemon by name.
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "pokemon/" + pokemonName
	fullUrl := pokeBaseURL + endpoint

	
	// Check the cache before doing a network call.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit for", fullUrl)
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	fmt.Println("Cache miss for", fullUrl)
	req, error := http.NewRequest("GET", fullUrl, nil)
	if error != nil {
		return Pokemon{}, error
	}

	res, error := c.httpClient.Do(req)
	if error != nil {
		return Pokemon{}, error
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullUrl, data)
	return pokemon, nil
}
