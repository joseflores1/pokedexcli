package main

import (
	"fmt"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: explore command requires only the <location_area> argument")
	}
	resources, err := pokeapi.GetNamedResources[pokeapi.PokemonByLocationArea](pokeapi.LocationAreaURL, args[0], config.Cache, config.Client)
	if err != nil {
		return fmt.Errorf("error getting Pokemon's names: %w", err)
	}

	fmt.Printf("Printing Pokemon's names for the %s location!\n\n", args[0])
	fmt.Print("Pokemon's names:\n\n")	
	for i, resource := range resources.PokemonEncounters {
		fmt.Printf("%d. %s\n", i + 1, resource.Pokemon.Name)
	}
	return nil
}