package commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var TheCommands = map[string]cliCommand{}

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
	}
}
