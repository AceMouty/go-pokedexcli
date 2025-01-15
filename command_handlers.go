package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

func handleExitCommand(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func handleHelpCommand(cfg *Config, args ...string) error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, command := range getCommands() {
		message := fmt.Sprintf("%v: %v", command.name, command.description)
		fmt.Println(message)
	}
	fmt.Print("\n\n")
	return nil
}

func handleMapCommand(cfg *Config, args ...string) error {
	requestUrl := baseUrl

	if cfg.next != nil {
		requestUrl = *cfg.next
	}

	locationApiResponse, err := cfg.pokeApiClient.GetLocationAreas(&requestUrl)

	if err != nil {
		return err
	}

	cfg.next = locationApiResponse.Next
	cfg.previous = locationApiResponse.Previous

	fmt.Println("Found areas")
	for _, location := range locationApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func handleMapbCommand(cfg *Config, args ...string) error {
	if cfg.previous == nil {
		return nil
	}

	locationApiResponse, err := cfg.pokeApiClient.GetLocationAreas(cfg.previous)

	if err != nil {
		return err
	}

	cfg.next = locationApiResponse.Next
	cfg.previous = locationApiResponse.Previous

	for _, location := range locationApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func handleExploreCommand(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	reqUrl := baseUrl + name
	location, err := cfg.pokeApiClient.GetLocationArea(&reqUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func handleCatchCommand(cfg *Config, args ...string) error {
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	pokemon, err := cfg.pokeApiClient.GetPokemon(&pokemonName)

	if err != nil {
		return err
	}

	maxCatchRate := 100
	catchRate := maxCatchRate - pokemon.BaseExperience/5
	randomNum := rand.Intn(maxCatchRate)
	catchPokemon := randomNum < catchRate

	if catchPokemon {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}

func handleInspectCommand(cfg *Config, args ...string) error {

	pokemonName := args[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("no")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}

func handlePokedexCommand(cfg *Config, args ...string) error {
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	return nil
}
