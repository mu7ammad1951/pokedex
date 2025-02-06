package main

import "fmt"

func commandHelp(cfg *config, arg *string) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil

}
