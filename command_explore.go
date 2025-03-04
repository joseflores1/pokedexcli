package main

import (
	"fmt"
)

func commandExplore(config *config, locationArea string) error {
	resources, err := config.Client.GetPokemonByLocationArea(locationArea, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting Pokemon's names: %w", err)
	}

	fmt.Print("Printing Pokemon's names!\n\n")
	fmt.Print("Pokemon's names:\n\n")	
	for i, resource := range resources.PokemonEncounters {
		fmt.Printf("%d. %s\n", i + 1, resource.Pokemon.Name)
	}
	return nil
}