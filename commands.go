package main

import (
	"fmt"
	"log"
	"os"
)

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
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
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

func commandMapB(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		fmt.Println("No previous location areas.")
		return nil
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
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
