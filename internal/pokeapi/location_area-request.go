package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaList(pageUrl *string) (LocationAreaResponse, error) {
	endpoint := "location-area/"
	fullUrl := pokeBaseURL + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	//check the cache
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
	req, error := http.NewRequest("GET", fullUrl, nil)
	if error != nil {
		return LocationAreaResponse{}, error
	}

	res, error := c.httpClient.Do(req)
	if error != nil {
		return LocationAreaResponse{}, error
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationAreaResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "location-area/" + locationAreaName
	fullUrl := pokeBaseURL + endpoint

	//check the cache
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
