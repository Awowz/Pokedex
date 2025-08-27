package main

import (
	"fmt"

	"github.com/Awowz/Pokedex/internal/pokeapi"
)

func commandExplore(conf *config, argz []string) {
	if len(argz) <= 0 {
		fmt.Printf("explore command must be followed by a area name ex: 'explore pastoria-city-area'\n")
	}
	fmt.Printf("Exploring %s...\n", argz[0])
	url := pokeapi.BASE_URL + pokeapi.URL_LOCATIONS + "/" + argz[0]
	//todo finilize logic, add to repl.go for user to utilize
}
