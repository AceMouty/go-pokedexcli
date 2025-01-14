package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleHelpCommand(paginationCfg *paginationConfig) error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, command := range getCommands() {
		message := fmt.Sprintf("%v: %v", command.name, command.description)
		fmt.Println(message)
	}
	fmt.Print("\n\n")

	return nil
}

func handleExitCommand(paginationCfg *paginationConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// endpoint: https://pokeapi.co/api/v2/location/
type Location struct {
	Name string `json:"name"`
}

type LocationApiResponse struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func handleMapCommand(paginationCfg *paginationConfig) error {
	requestUrl := "https://pokeapi.co/api/v2/location-area/"

	if paginationCfg.next != "" {
		requestUrl = paginationCfg.next
	}

	locationApiResponse, err := makeGetRequest(requestUrl)

	if err != nil {
		return err
	}

	paginationCfg.next = locationApiResponse.Next
	paginationCfg.previous = locationApiResponse.Previous

	for _, location := range locationApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func handleMapbCommand(paginationCfg *paginationConfig) error {
	if paginationCfg.previous == "" {
		return nil
	}

	locationApiResponse, err := makeGetRequest(paginationCfg.previous)

	if err != nil {
		return err
	}

	paginationCfg.next = locationApiResponse.Next
	paginationCfg.previous = locationApiResponse.Previous

	for _, location := range locationApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func makeGetRequest(requestUrl string) (LocationApiResponse, error) {
	var locationApiResponse LocationApiResponse

	resp, err := http.Get(requestUrl)
	if err != nil {
		return locationApiResponse, err
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationApiResponse, err
	}

	err = json.Unmarshal(jsonData, &locationApiResponse)
	if err != nil {
		return locationApiResponse, err
	}

	return locationApiResponse, nil
}
