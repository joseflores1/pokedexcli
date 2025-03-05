package main

import (
	"fmt"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

func commandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: inspect command requires only the <pokemon_name> argument")
	}
	if _, ok := config.Pokedex.PokemonList[args[0]]; !ok {
		fmt.Printf("You have not caught a %s yet!\n\n", args[0])
		fmt.Println("Your Pokedex:")
		printPokedex(config.Pokedex.PokemonList)
		return nil
	}

	pokemon := config.Pokedex.PokemonList[args[0]]
	printPokemonInfo(pokemon)

	return nil
}

func printPokemonInfo(pokemon pokeapi.PokemonByName) {
	statList := ""
	for _, stat := range pokemon.Stats {
		statList += fmt.Sprintf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	typeList := ""
	for _, typeInfo := range pokemon.Types {
		typeList += fmt.Sprintf("\t- %s\n", typeInfo.Type.Name)
	}
	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
Stats:
%s
Types:
%s
`, pokemon.Name, pokemon.Height, pokemon.Weight, statList, typeList)
}