package main

import (
	"time"

	"github.com/Awowz/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{pokeapiClient: pokeClient, pokeDex: make(map[string]pokeapi.PokemonData)}
	startRepl(cfg)
}
