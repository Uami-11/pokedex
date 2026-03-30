package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
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

var (
	TheCommands = map[string]cliCommand{}
	TheConfig   config
)

func init() {
	TheCommands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func() error {
				return Help(TheCommands)
			},
		},
		"map": {
			Name:        "map",
			Description: "Shows the next 20 locations",
			Callback: func() error {
				return Map(&TheConfig)
			},
		},
		"mapb": {
			Name:        "map back",
			Description: "Shows the previous 20 locations",
			Callback: func() error {
				return MapBack(&TheConfig)
			},
		},
	}
}

func fetchLocations(url string) (LocationAreaResponse, error) {
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
