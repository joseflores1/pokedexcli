package main

import (
	"fmt"
)
func commandMap(config *config, args ...string) error {
	resources, err := config.Client.GetUnnamedResources(config.Next, config.Endpoint, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting Resources's names: %w", err)
	}

	fmt.Print("Printing next Resources's names!\n\n")
	fmt.Print("Resources's names:\n\n")	
	for i, resource := range resources.Results {
		fmt.Printf("%d. %s\n", i + 1, resource.Name)
	}
	config.Previous = resources.Previous
	config.Next = resources.Next
	return nil
}

func commandMapb(config *config, args ...string) error {
	if config.Previous == nil {
		return fmt.Errorf("error because nil page (likely got to the last \"previous\" page)")
	}

	resources, err := config.Client.GetUnnamedResources(config.Previous, config.Endpoint, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting resources's names: %w", err)
	}

	config.Previous = resources.Previous
	config.Next = resources.Next

	fmt.Print("Printing previous resources names!\n\n")
	fmt.Print("Resources's names:\n\n")
	for i, resource := range resources.Results {
		fmt.Printf("%d, %s\n", i + 1, resource.Name)
	}

	return nil
}
