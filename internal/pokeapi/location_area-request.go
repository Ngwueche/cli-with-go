package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocationAreaList fetches a paginated list of location areas.
func (c *Client) GetLocationAreaList(pageUrl *string) (LocationAreaResponse, error) {
	endpoint := "location-area/"
	fullUrl := pokeBaseURL + endpoint

	// If a page URL is provided, use it instead of the default endpoint.
	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	// Check the cache before doing a network call.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit for", fullUrl)
		locationAreaResponse := LocationAreaResponse{}
		err := json.Unmarshal(dat, &locationAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaResponse, nil
	}
	fmt.Println("Cache miss for", fullUrl)
	// http.NewRequest builds a request without executing it.
	req, error := http.NewRequest("GET", fullUrl, nil)
	if error != nil {
		return LocationAreaResponse{}, error
	}

	// Do sends the request and returns a response or error.
	res, error := c.httpClient.Do(req)
	if error != nil {
		return LocationAreaResponse{}, error
	}

	// Ensure the response body is closed when the function returns.
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	// ReadAll consumes the body into a byte slice.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Unmarshal JSON into a struct.
	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	// Store raw response bytes in cache for next time.
	c.cache.Add(fullUrl, data)
	return locationAreaResponse, nil
}

// GetLocationArea fetches details for a single location area by name.
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "location-area/" + locationAreaName
	fullUrl := pokeBaseURL + endpoint

	// Check the cache before doing a network call.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit for", fullUrl)
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("Cache miss for", fullUrl)
	req, error := http.NewRequest("GET", fullUrl, nil)
	if error != nil {
		return LocationArea{}, error
	}

	res, error := c.httpClient.Do(req)
	if error != nil {
		return LocationArea{}, error
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationArea, nil
}
