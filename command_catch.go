package main

import (
	"fmt"
)

func commandCatch(c *config, argz []string) error {
	if len(argz) <= 0 {
		return fmt.Errorf("please provide a pokemons name youd like to catch")
	}
	pokeDetails, err := c.pokeapiClient.GetPokemonDetails(argz[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", argz[0])
	if pokeDetails.AttemptCatch() {
		fmt.Printf("%s was caught!\n", pokeDetails.Name)
		c.pokeDex[pokeDetails.Name] = pokeDetails
	} else {
		fmt.Printf("%s escaped!\n", pokeDetails.Name)
	}
	return nil
}
