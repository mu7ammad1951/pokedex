package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mu7ammad1951/pokedex/internal/pokeapi"
)

type config struct {
	PokeAPIClient   pokeapi.Client
	NextLocationURL *string
	PrevLocationURL *string
}
type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}
		commandName := cleanedText[0]
		var arg *string
		if len(cleanedText) > 1 {
			arg = &cleanedText[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	splitStrings := strings.Fields(lowercaseText)

	return splitStrings
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Retrieves and displays Pokémon location areas from the next available page in the Pokédex API",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Retrieves and displays Pokémon location areas from the previous available page in the Pokédex API.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon in a given location area",
			callback:    commandExplore,
		},
	}
}
