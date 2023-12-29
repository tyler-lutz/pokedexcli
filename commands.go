package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error {
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

func commandMapB(cfg *config, args ...string) error {
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println(args)
		return errors.New("no location area provided")
	}

	name := args[0]
	res, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + res.Name + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// random chance of catching pokemon based on base experience
	chanceToEscape := rand.Intn(pokemon.BaseExperience)

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	if chanceToEscape > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeEntry := range pokemon.Types {
		fmt.Printf("- %s\n", typeEntry.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
