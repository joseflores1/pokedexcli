package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config *Config = &Config{
	Next: &locationEndpoint,
	Previous: nil,
}

func startRepl() error {
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
		fmt.Println("------------------------------------------------------")

		if err != nil {
			return fmt.Errorf("error when trying to use \"%s\" command\n: %w", command, err)
		}
	}

}
func cleanInput(text string) (words []string) {
	lowerTrim := strings.Trim(strings.ToLower(text), " ")
	wordsSlice := strings.Split(lowerTrim, " ")
	return wordsSlice

}


