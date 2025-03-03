package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetResources(pageURL *string) (UnnamedResources, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UnnamedResources{}, fmt.Errorf("error making GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return UnnamedResources{}, fmt.Errorf("error getting response: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return UnnamedResources{}, fmt.Errorf("error: StatusCode = %d", res.StatusCode) 
	}	

	var resources UnnamedResources
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		return UnnamedResources{}, fmt.Errorf("error decoding response: %w", err)
	}

	return resources, nil
}


