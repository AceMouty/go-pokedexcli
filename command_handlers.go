package main

import (
	"fmt"
	"os"
)

func handleHelpCommand(cfg *Config) error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, command := range getCommands() {
		message := fmt.Sprintf("%v: %v", command.name, command.description)
		fmt.Println(message)
	}
	fmt.Print("\n\n")

	return nil
}

func handleExitCommand(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// endpoint: https://pokeapi.co/api/v2/location/

func handleMapCommand(cfg *Config) error {
	requestUrl := "https://pokeapi.co/api/v2/location-area/"

	if cfg.next != nil {
		requestUrl = *cfg.next
	}

	locationApiResponse, err := cfg.pokeApiClient.GetLocationAreas(&requestUrl)

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

func handleMapbCommand(cfg *Config) error {
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
