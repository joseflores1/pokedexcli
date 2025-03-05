package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
	"github.com/joseflores1/pokedexcli/internal/pokecache"
)

type config struct {
	Client *pokeapi.Client
	Next *string 
	Previous *string 
	Endpoint string
	Cache *pokecache.Cache
	Pokedex *pokeapi.Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays in each subsequent call the next 20 location's names",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays in each subsequent call the previous 20 location's names",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Given an input location area, displays all the pokemon in it",
			callback: commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Catches <pokemon_name> with probability defined by its base experience value",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Displays information of <pokemon_name> only if caught" ,
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays the caught pokemon saved in the pokedex",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) (words []string) {
	lowerTrim := strings.Trim(strings.ToLower(text), " ")
	wordsSlice := strings.Split(lowerTrim, " ")
	return wordsSlice
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)	

	selectionMenu(config, *scanner)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		
		if len(text) == 0 {
			continue
		}

		command := cleanText[0]
		args := []string{}

		if len(cleanText) > 1 {
			args = cleanText[1:]
		}

		commandValue, ok := getCommands()[command]
		if !ok {
			fmt.Println("------------------------------------------------------")
			fmt.Printf("Unknown command\n")
			fmt.Println("------------------------------------------------------")
			continue
		}

		fmt.Println("------------------------------------------------------")
		err := commandValue.callback(config, args...)
		if err == nil {
			fmt.Println("------------------------------------------------------")
		}

		if err != nil {
			fmt.Printf("error when trying to use \"%s\" command\n%s", commandValue.name, err)
			fmt.Println("\n------------------------------------------------------")
		}
	}

}

func getEndpoints()map[string]string {
	return map[string]string{
		"1": "/location-area",
		"2": "/location",
	}
}

func selectionMenu(config *config, scanner bufio.Scanner) {
	fmt.Printf(`
----------------------------------------------------
Input a number to select the endpoint for retrieval:

1. %s
2. %s
----------------------------------------------------  
`, getEndpoints()["1"][1:], getEndpoints()["2"][1:])

	var path string

	for {
		fmt.Print("Enter your option: ")
		scanner.Scan()
		option:= scanner.Text()
		switch option {
		case "1":
			path = getEndpoints()[option]
		case "2":
			path = getEndpoints()[option]
		default:
			fmt.Println("----------------------")	
			fmt.Println("Input a valid number: ")
			fmt.Println("----------------------")	
		}
		if path != "" {
			break
		}
	}
	config.Endpoint = path
	fmt.Println("----------------------")	
	fmt.Printf("You chose the %s option!\n", path[1:])
	fmt.Println("----------------------")	

}
