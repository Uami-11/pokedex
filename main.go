package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Uami-11/pokedex/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokédex > ")
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Failed to take user input")
		}

		userInput := scanner.Text()
		userWords := cleanInput(userInput)
		userCommand := userWords[0]
		var validCommand bool

		for command := range commands.TheCommands {
			if userCommand == command {
				validCommand = true
				err := commands.TheCommands[command].Callback()
				if err != nil {
					fmt.Println("%w", err)
				}
			}
		}

		if !validCommand {
			fmt.Println("Unkown command")
		}
	}
}
