// Package commands...
package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	pokecache "github.com/Uami-11/pokedex/internal"
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

func Explore(location string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	val, ok := pokecache.LocationCache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		pokecache.LocationCache.Add(url, val)
	}

	var data LocationInfo
	if err := json.Unmarshal(val, &data); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")

	for _, enc := range data.PokemonEncounters {
		fmt.Printf("- %s\n", enc.ThePokemon.Name)
	}

	return nil
}
