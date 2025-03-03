package  main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}


type Config struct {
	Next *string `json:"next"`
	Previous *string `json:"previous"`
}

var locationEndpoint string = "https://pokeapi.co/api/v2/location-area/"

func getLocations(url *string) (LocationAreas, error) {
	if url == nil {
		return LocationAreas{}, fmt.Errorf("error because of nil url (you likely arrived at the last \"previous\" url)")
	}
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error making GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error getting response: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("error: StatusCode = %d", res.StatusCode) 
	}	

	var locations LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error decoding response: %w", err)
	}
	return locations, nil
}


