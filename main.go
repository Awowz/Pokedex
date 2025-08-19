package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedUserInputArray := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", cleanedUserInputArray[0])
	}
}
