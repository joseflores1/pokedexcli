package main

import (
	"time"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
	"github.com/joseflores1/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config:= &config{
		Client: pokeClient,
		Next: nil,
		Previous: nil,
		Endpoint: "",
		Cache: pokecache.NewCache(time.Minute * 5),
	}

	startRepl(config)
}