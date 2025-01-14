package main

import (
	"bufio"
	"fmt"
	"strings"
)

func startRepl(s *bufio.Scanner, paginationCfg *paginationConfig) {

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
		command, ok := getCommands()[commandname]
		if !ok {
			fmt.Println("Unkown command")
			continue
		}

		command.callback(paginationCfg)
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
	callback    func(*paginationConfig) error
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

	return commands
}