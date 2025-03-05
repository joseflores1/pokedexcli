package main

import (
	"fmt"
	"math/rand"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

var r rand.Rand = *rand.New(rand.NewSource(0))

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: explore command requires only the <pokemon_name> argument")
	}
	if _, ok := config.Pokedex.PokemonList[args[0]]; ok {
		fmt.Printf("%s was already caught!\n\n", args[0])
		commandPokedex(config)
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
		fmt.Println("You may now inspect it with the \"inspect\" command")
		config.Pokedex.PokemonList[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

