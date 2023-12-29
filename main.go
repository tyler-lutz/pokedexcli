package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tyler-lutz/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the pokemon world.",
			callback:    commandMap,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	pokeapiClient := pokeapi.NewClient()

	res, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}

func parseInput(input string) []string {
	lowered := strings.ToLower(input)
	return strings.Fields(lowered)
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}
		parsedInput := parseInput(input)
		userCommand := parsedInput[0]
		if command, ok := getCommands()[userCommand]; ok {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", userCommand)
			continue
		}
	}
}
