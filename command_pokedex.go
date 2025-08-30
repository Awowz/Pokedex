package main

import "fmt"

func commandPokedex(c *config, argz []string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range c.pokeDex {
		fmt.Printf(" - %v\n", val.Name)
	}
	return nil
}
