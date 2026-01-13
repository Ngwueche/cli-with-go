package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaList() (LocationAreaResponse, error) {
	endpoint := "location-area/"
	fullUrl := pokeBaseURL + endpoint

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
	
	return locationAreaResponse, nil
}
