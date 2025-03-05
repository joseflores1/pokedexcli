package pokeapi

import (
)

type unnamedResources struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


type PokemonByLocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}


type PokemonByName struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
}

type Pokedex struct {
	PokemonList map[string]PokemonByName
}

type namedResource interface {
	PokemonByLocationArea | PokemonByName
}
