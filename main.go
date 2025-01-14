package main

import (
	"bufio"
	"github.com/acemouty/pokedexcli/internal/pokeapi"
	"os"
	"time"
)

type Config struct {
	pokeApiClient pokeapi.Client
	next          *string
	previous      *string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	apiClient := pokeapi.NewClient(5 * time.Second)
	cfg := Config{pokeApiClient: apiClient}

	startRepl(scanner, &cfg)
}
