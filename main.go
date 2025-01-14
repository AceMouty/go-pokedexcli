package main

import (
	"bufio"
	"os"
)

type paginationConfig struct {
	next     string
	previous string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	paginationCfg := paginationConfig{next: "", previous: ""}

	startRepl(scanner, &paginationCfg)
}
