package main

import (
	"fmt"
	"sort"
)
func commandPokedex(config *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("error: pokedex command accepts no arguments")
	}
	var keys []string
	for k := range config.Pokedex.PokemonList {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println("Your Pokedex:")

	for i, k := range keys {
		fmt.Printf("%d. %s\n", i + 1, config.Pokedex.PokemonList[k].Name)
	}
	return nil
}