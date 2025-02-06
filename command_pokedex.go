package main

import "fmt"

func commandPokedex(cfg *config, arg *string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %s", pokemon.Name)
	}
	return nil
}
