package main

import (
	"time"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config:= &config{
		Client: pokeClient,
		Next: nil,
		Previous: nil,
	}

	startRepl(config)
}