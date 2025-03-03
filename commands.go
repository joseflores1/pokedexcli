package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
	}
}


func commandExit(config *Config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return fmt.Errorf("error when printing goodbye message: %w", err)
	}
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
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

func commandMap(config *Config) error {
	fmt.Print("Printing next location's names!\n\n")
	locations, err := getLocations(config.Next)
	if err != nil {
		return fmt.Errorf("error getting location's names: %w", err)
	}
	fmt.Print("Location's names:\n\n")	
	for i, location := range locations.Results {
		fmt.Printf("%d. %s\n", i + 1, location.Name)
	}
	config.Previous = config.Next
	config.Next = locations.Next
	return nil
}

func commandMapb(config *Config) error {
	locations, err := getLocations(config.Previous)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error getting location's names (first change): %w", err)
	}

	config.Next = config.Previous
	config.Previous = locations.Previous

	locations, err = getLocations(config.Previous)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error getting location's names (second change): %w", err)
	}

	fmt.Print("Printing previous location's names!\n\n")
	fmt.Print("Location's names:\n\n")
	for i, location := range locations.Results {
		fmt.Printf("%d, %s\n", i + 1, location.Name)
	}

	return nil
}