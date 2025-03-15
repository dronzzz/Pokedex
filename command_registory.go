package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Location  `json:"results"`
}

var commands map[string]cliCommand
var locationResponse LocationResponse

func fetchCommands() {

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    listAllCommands,
		},
		"map": {
			name:        "map",
			description: "displays the next 20 locations",
			callback:    fetchMapNext,
		},
		"mapb": {
			name:        "map",
			description: "displays the previous 20 locations",
			callback:    fetchMapPrev,
		},
	}

}

func listAllCommands() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commands {
		fmt.Printf("\n%s : %s", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func fetchMapNext() error {
	godotenv.Load(".env")
	BASE_URL := os.Getenv("BASE_URL")
	var url string

	if locationResponse.Next != "" {
		url = locationResponse.Next
	} else {
		url = BASE_URL + "/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err := json.Unmarshal(body, &locationResponse); err != nil {
		log.Fatal(err)
	}
	var location []Location = locationResponse.Results
	fmt.Println()

	for _, loc := range location {
		fmt.Println(loc.Name)
	}

	return nil
}

func fetchMapPrev() error {

	if locationResponse.Previous == nil {
		fmt.Println(" This is the first page")
		return nil
	}
	// prevURL, ok := locationResponse.Previous.(string)
	// if !ok {
	// 	fmt.Println("Error: Previous URL is not a valid string")
	// 	return nil
	// }

	res, err := http.Get(locationResponse.Previous.(string))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := json.Unmarshal(body, &locationResponse); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
