package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

var r rand.Rand = *rand.New(rand.NewSource(0))

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: explore command requires only the <pokemon_name> argument")
	}
	if _, ok := config.Pokedex.PokemonList[args[0]]; ok {
		fmt.Printf("%s was already caught!\n\n", args[0])
		fmt.Println("Your Pokedex:")
		printPokedex(config.Pokedex.PokemonList)
		return nil
	}
	pokemon, err := pokeapi.GetNamedResources[pokeapi.PokemonByName](pokeapi.PokemonURL, args[0], config.Cache, config.Client)
	if err != nil {
		return fmt.Errorf("error getting pokemon's info: %w", err)
	}

	res := float64(r.Intn(pokemon.BaseExperience)) 
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res < 70 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Pokedex.PokemonList[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func printPokedex(pokedex map[string]pokeapi.PokemonByName) {
	var keys []string

	for k := range pokedex {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i, k := range keys {
		fmt.Printf("%d. %s\n", i + 1, pokedex[k].Name)
	}
}