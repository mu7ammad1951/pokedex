package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, arg *string) error {
	url := cfg.NextLocationURL
	resAreas, err := cfg.PokeAPIClient.ListLocations(url)
	if err != nil {
		return fmt.Errorf("error retrieving locations")
	}

	cfg.NextLocationURL = resAreas.Next
	cfg.PrevLocationURL = resAreas.Previous

	fmt.Println()
	for _, location := range resAreas.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
	return nil
}

func commandMapb(cfg *config, arg *string) error {
	if cfg.PrevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	url := cfg.PrevLocationURL
	resAreas, err := cfg.PokeAPIClient.ListLocations(url)
	if err != nil {
		return fmt.Errorf("error retrieving locations")
	}

	cfg.NextLocationURL = resAreas.Next
	cfg.PrevLocationURL = resAreas.Previous

	fmt.Println()
	for _, location := range resAreas.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
	return nil
}
