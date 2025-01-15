package main

import (
	"bufio"
	"fmt"
	"strings"
)

func startRepl(s *bufio.Scanner, cfg *Config) {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()

	for {
		fmt.Print("Pokedex > ")

		input, err := getUserInput(s)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		input_slice := cleanInput(input)
		if len(input_slice) == 0 {
			fmt.Println()
			continue
		}

		commandname := input_slice[0]
		args := []string{}
		if len(input_slice) > 1 {
			args = input_slice[1:]
		}

		command, ok := getCommands()[commandname]
		if !ok {
			fmt.Println("Unkown command")
			continue
		}

		command.callback(cfg, args...)
	}
}

func getUserInput(s *bufio.Scanner) (string, error) {
	s.Scan()
	input := s.Text()

	if err := s.Err(); err != nil {
		return "", err
	}

	return input, nil
}

func cleanInput(text string) []string {

	trimmed_string := strings.TrimSpace(text)
	words := strings.Fields(trimmed_string)

	for idx := range words {
		word := words[idx]
		words[idx] = strings.ToLower(word)
	}

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type commandMap = map[string]cliCommand

func getCommands() commandMap {
	commands := make(commandMap)

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    handleExitCommand,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    handleHelpCommand,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "displays next 20 locations",
		callback:    handleMapCommand,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "displays previous 20 locations",
		callback:    handleMapbCommand,
	}

	commands["explore"] = cliCommand{
		name:        "explore <location_name>",
		description: "Explore a location",
		callback:    handleExploreCommand,
	}

	commands["catch"] = cliCommand{
		name:        "catch <pokemon_name>",
		description: "Attempt to catch a pokemon",
		callback:    handleCatchCommand,
	}

	commands["inspect"] = cliCommand{
		name:        "inspect <pokemon_name>",
		description: "Display information about a specific pokemon",
		callback:    handleInspectCommand,
	}

	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Display names of all pokemon caught",
		callback:    handlePokedexCommand,
	}

	return commands
}
