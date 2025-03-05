package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, args ...string) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return fmt.Errorf("error when printing goodbye message: %w", err)
	}
	os.Exit(0)
	return nil
}