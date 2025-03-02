package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
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
		err := commandValue.callback()
		fmt.Println("------------------------------------------------------")

		if err != nil {
			fmt.Printf("error when trying to use \"%s\" command\n:", command)
		}
	}

}
func cleanInput(text string) (words []string) {
	lowerTrim := strings.Trim(strings.ToLower(text), " ")
	wordsSlice := strings.Split(lowerTrim, " ")
	return wordsSlice

}


