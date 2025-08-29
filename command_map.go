package main

import (
	"fmt"

	"github.com/Awowz/Pokedex/internal/pokeapi"
)

func commandMapf(conf *config, argz []string) error {
	pokeMapData, err := conf.pokeapiClient.ShallowListLocations(conf.next)
	if err != nil {
		return err
	}

	setandPrintMap(conf, pokeMapData)

	return nil
}

func commandMapb(conf *config, argz []string) error {
	pokeMapData, err := conf.pokeapiClient.ShallowListLocations(conf.previous)
	if err != nil {
		return err
	}

	setandPrintMap(conf, pokeMapData)
	return nil
}

func setandPrintMap(conf *config, pokeMapData pokeapi.PokeMap) {
	conf.next = pokeMapData.Next
	conf.previous = pokeMapData.Previous

	for _, x := range pokeMapData.Results {
		fmt.Println(x.Name)
	}
}
