package main

import (
	"fmt"

	"github.com/Awowz/Pokedex/internal/pokeapi"
)

func commandExplore(conf *config, argz []string) error {
	if len(argz) <= 0 {
		return fmt.Errorf("explore command must be followed by a area name ex: 'explore pastoria-city-area'")
	}
	fmt.Printf("Exploring %s...\n", argz[0])
	url := pokeapi.BASE_URL + pokeapi.URL_LOCATIONS + "/" + argz[0]
	pokeEcnounters, err := conf.pokeapiClient.ListPokemonFromLocation(url)
	if err != nil {
		return err
	}
	pokeEcnounters.DisplayPokemon()
	return nil
}
