package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}


func commandExit() error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return fmt.Errorf("error when printing goodbye message: %w", err)
	}
	os.Exit(0)
	return nil
}

func commandHelp() error {

	_, err := fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	if err != nil {
		return fmt.Errorf("error when printing help message: %w", err)
	}
	for _, v := range getCommands() {
		_, err := fmt.Printf("%s: %s\n", v.name, v.description)
		if err != nil {
			return fmt.Errorf("error when printing command description list: %w", err)
		}
	}
	return nil
}