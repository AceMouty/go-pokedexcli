package main

import (
	"errors"
	"fmt"
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
