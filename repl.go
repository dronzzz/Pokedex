package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		value, exists := commands[input]

		if exists {
			if value.callback() != nil { //bad practice should be err:= value.callback() ; err!= nil
				fmt.Printf("Error while calling %v /t", input)
			}

		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := []string{}
	j := 0
	for i := 0; i < len(text); i++ {
		if i == len(text) || text[i] == ' ' {
			if i > j {
				output = append(output, strings.ToLower(text[j:i]))
			}
			j = i + 1
		}
	}

	return output

}
