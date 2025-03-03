package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joseflores1/pokedexcli/internal/pokeapi"
)

type config struct {
	Client pokeapi.Client
	Next *string 
	Previous *string 
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		
		if len(text) == 0 {
			continue
		}

		command := cleanText[0]


		commandValue, ok := getCommands()[command]
		if !ok {
			fmt.Println("------------------------------------------------------")
			fmt.Printf("Unknown command\n")
			fmt.Println("------------------------------------------------------")
			continue
		}

		fmt.Println("------------------------------------------------------")
		err := commandValue.callback(config)
		if err == nil {
			fmt.Println("------------------------------------------------------")
		}

		if err != nil {
			fmt.Printf("error when trying to use \"%s\" command\n%s", command, err)
			fmt.Println("\n------------------------------------------------------")
		}
	}

}
