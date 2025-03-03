package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joseflores1/pokedexcli/internal/pokecache"
)

func (c *Client) GetResources(pageURL *string, endpoint string, cache *pokecache.Cache) (UnnamedResources, error) {
	url := baseURL + endpoint + "?offset=0&limit=20"
	if pageURL != nil {
		url  = *pageURL
	}
	if value, ok := cache.Get(url); ok {
		fmt.Println("Retrieving data from cache!")
		resources := UnnamedResources{}
		err := json.Unmarshal(value, &resources)
		if err != nil {
			return UnnamedResources{}, fmt.Errorf("error unmarshalling data from cache: %w", err)
		}
		return resources, nil

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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return UnnamedResources{}, fmt.Errorf("error reading response data: %w", err)
	}
	cache.Add(url, data)

	fmt.Println("Retrieving data from request!")
	resources := UnnamedResources{}
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return UnnamedResources{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resources, nil
}


