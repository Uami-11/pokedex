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
	fmt.Print("         .-. \\_/ .-.\n" +
		"         \\.-\\/=\\/.-/\n" +
		"      '-./___|=|___\\.-'\n" +
		"     .--| \\|/`\"\"`\\|/ |--.\n" +
		"    (((_)\\  .---.  /(_)))\n" +
		"     `\\ \\_`-.   .-'_/ /`_\n" +
		"       '.__       __.'(_))\n" +
		"           /     \\     //\n" + "          |       |__.'/\n" +
		"          \\       /--'`\n" +
		"      .--,-' .--. '----.\n" +
		"     '----`--'  '--`----'\n")
	for {
		fmt.Print("Pokédex > ")
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Failed to take user input")
		}

		userInput := scanner.Text()
		userWords := cleanInput(userInput)
		userCommand := userWords[0]
		userArguments := userWords[1:]
		var validCommand bool

		for command := range commands.TheCommands {
			if userCommand == command {
				validCommand = true
				err := commands.TheCommands[command].Callback(userArguments)
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
