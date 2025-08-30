package main

import (
	"fmt"
)

func commandInspect(c *config, argz []string) error {
	if len(argz) <= 0 {
		return fmt.Errorf("please include the pokemon you want to inspect")
	}
	thePokemon, valid := c.pokeDex[argz[0]]
	if !valid {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf(`Height: %v
Weight: %v
Stats:
`, thePokemon.Height, thePokemon.Weight)
	for _, stat := range thePokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.Base_stat)
	}
	fmt.Printf("Types:\n")
	for _, val := range thePokemon.Types {
		fmt.Printf("  - %v\n", val.Type.Name)
	}
	return nil
}
