package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joseflores1/pokedexcli/internal/pokecache"
)

func (c *Client) GetUnnamedResources(pageURL *string, endpoint string, cache *pokecache.Cache) (unnamedResources, error) {
	url := BaseURL + endpoint + "?offset=0&limit=20"
	if pageURL != nil {
		url  = *pageURL
	}
	if value, ok := cache.Get(url); ok {
		fmt.Println("Retrieving data from cache!")
		resources := unnamedResources{}
		err := json.Unmarshal(value, &resources)
		if err != nil {
			return unnamedResources{}, fmt.Errorf("error unmarshalling data from cache: %w", err)
		}
		return resources, nil

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return unnamedResources{}, fmt.Errorf("error making GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return unnamedResources{}, fmt.Errorf("error getting response: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return unnamedResources{}, fmt.Errorf("error: StatusCode = %d", res.StatusCode) 
	}	

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return unnamedResources{}, fmt.Errorf("error reading response data: %w", err)
	}
	cache.Add(url, data)

	fmt.Println("Retrieving data from request!")
	resources := unnamedResources{}
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return unnamedResources{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resources, nil
}


func GetNamedResources[T namedResource](baseURL, name string, cache *pokecache.Cache, client *Client) (T, error) {
	url := baseURL + name

	var empty T
	if value, ok := cache.Get(url); ok {
		fmt.Println("Retrieving data from cache!")
		resources := empty
		err := json.Unmarshal(value, &resources)
		if err != nil {
			return empty, fmt.Errorf("error unmarshalling data from cache: %w", err)
		}
		return resources, nil

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return empty, fmt.Errorf("error making GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return empty, fmt.Errorf("error getting response: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return empty, fmt.Errorf("error: StatusCode = %d", res.StatusCode) 
	}	

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return empty, fmt.Errorf("error reading response data: %w", err)
	}
	cache.Add(url, data)

	fmt.Println("Retrieving data from request!")
	resources := empty
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return empty, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resources, nil
}

