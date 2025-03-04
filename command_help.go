package main

import (
	"fmt"
	"sort"
)
func commandHelp(config *config, parameter string) error {
	_, err := fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	if err != nil {
		return fmt.Errorf("error when printing help message: %w", err)
	}

	allCommands := getCommands()

	keys := make([]string, 0, len(allCommands))

	for k := range allCommands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		_, err := fmt.Printf("%s: %s\n", allCommands[k].name, allCommands[k].description)
		if err != nil {
			return fmt.Errorf("error when printing command description list: %w", err)
		}
	}
	return nil
}