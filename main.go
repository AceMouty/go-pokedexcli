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
	pokedex       Pokedex
}

type Pokedex = map[string]pokeapi.Pokemon

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	httpTimeoutDuration := 5 * time.Second
	cacheInterval := 5 * time.Minute

	apiClient := pokeapi.NewClient(httpTimeoutDuration, cacheInterval)
	cfg := Config{pokeApiClient: apiClient, pokedex: Pokedex{}}

	startRepl(scanner, &cfg)
}
