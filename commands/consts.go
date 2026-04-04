package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	pokecache "github.com/Uami-11/pokedex/internal"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func([]string) error
}

type config struct {
	Next     *string
	Previous *string
}

type LocationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type LocationInfo struct {
	PokemonEncounters []struct {
		ThePokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var (
	TheCommands = map[string]cliCommand{}
	TheConfig   config
)

func init() {
	TheCommands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback: func(args []string) error {
				return Exit()
			},
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func(args []string) error {
				return Help(TheCommands)
			},
		},
		"map": {
			Name:        "map",
			Description: "Shows the next 20 locations",
			Callback: func(args []string) error {
				return Map(&TheConfig)
			},
		},
		"mapb": {
			Name:        "map back",
			Description: "Shows the previous 20 locations",
			Callback: func(args []string) error {
				return MapBack(&TheConfig)
			},
		},
		"explore": {
			Name:        "explore",
			Description: "Shows all the pokemon encounters in a specific area",
			Callback: func(args []string) error {
				if len(args) < 1 {
					return errors.New("usage: explore <location-name>")
				}
				return Explore(args[0])
			},
		},
	}
}

func fetchLocations(url string) (LocationAreaResponse, error) {
	if _, exists := pokecache.LocationCache.CacheEntries[url]; exists {

		locations, cached := pokecache.LocationCache.Get(url)
		if !cached {
			return LocationAreaResponse{}, errors.New("could not cache from the cache")
		}

		var data LocationAreaResponse

		if err := json.Unmarshal(locations, &data); err != nil {
			return data, nil
		} else {
			return data, err
		}

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	var data LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func Map(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if conf.Next != nil {
		url = *conf.Next
	}

	data, err := fetchLocations(url)
	if err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}

func MapBack(conf *config) error {
	if conf.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *conf.Previous

	data, err := fetchLocations(url)
	if err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}
