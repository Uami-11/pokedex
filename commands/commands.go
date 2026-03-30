// Package commands...
package commands

import (
	"fmt"
	"os"
)

func Exit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)

	return nil
}

func Help(theComms map[string]cliCommand) error {
	fmt.Print("Usage\n\n")
	for command := range theComms {
		fmt.Printf("%s: %s\n", command, theComms[command].Description)
	}

	return nil
}
