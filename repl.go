package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Awowz/Pokedex/internal/pokeapi"
)

// type struct for all commands
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	allCLIcommands := getAllCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedUserInputArray := cleanInput(scanner.Text())

		if len(cleanedUserInputArray) <= 0 {
			continue
		}

		userCommand, valid := allCLIcommands[cleanedUserInputArray[0]]
		if valid {
			err := userCommand.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// returns all commands and associated callback function
func getAllCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints all avaliable commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the name of 20 location areas in pokemon. each subsequent call to map will display the next 20 locations and so on",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in pokemon. each subsequent call to mapb will display the previous 20 locations and so on",
			callback:    commandMapb,
		},
	}

}

// This func would have benifited from "strings" packaged. i just wanted to exercise my brain
func cleanInput(text string) []string {
	var stringArray []string
	wordSetup := ""
	for index, char := range text {
		if char == ' ' {
			if wordSetup != "" {
				stringArray = append(stringArray, wordSetup)
				wordSetup = ""
				continue
			}
		} else {
			if char >= 'A' && char <= 'Z' {
				wordSetup += string(char + ('a' - 'A'))
			} else {
				wordSetup += string(char)
			}

			if len(text)-1 == index {
				stringArray = append(stringArray, wordSetup)
				wordSetup = ""
				continue
			}
		}
	}
	return stringArray
}
