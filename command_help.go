package main

import (
	"fmt"
)

func commandHelp(conf *config, argz []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, value := range getAllCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
