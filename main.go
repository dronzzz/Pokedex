package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to the Pokedex cliTool!")
	fmt.Println("Type 'help' to get assistance or 'exit' to exit the program ")
	fetchCommands()
	Repl()
}
