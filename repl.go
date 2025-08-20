package main

import (
	"bufio"
	"fmt"
	"os"
)

// type struct for all commands
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	allCLIcommands := getAllCommands()
	mapconfig := config{next: nil, previous: nil}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedUserInputArray := cleanInput(scanner.Text())

		if len(cleanedUserInputArray) <= 0 {
			continue
		}

		userCommand, valid := allCLIcommands[cleanedUserInputArray[0]]
		if valid {
			err := userCommand.callback(&mapconfig)
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
			callback:    commandMap,
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
