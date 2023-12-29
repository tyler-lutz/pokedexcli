package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tyler-lutz/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the pokemon world.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of pokemon that can be found in a given area",
			callback:    commandExplore,
		},
	}
}

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func parseInput(input string) []string {
	lowered := strings.ToLower(input)
	return strings.Fields(lowered)
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
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
		args := []string{}
		if len(parsedInput) > 1 {
			args = parsedInput[1:]
		}
		if command, ok := getCommands()[userCommand]; ok {
			err := command.callback(&cfg, args...)
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
