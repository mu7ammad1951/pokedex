package main

import (
	"fmt"

	"github.com/mu7ammad1951/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, location *string) error {
	locationInfo, err := cfg.PokeAPIClient.ExploreLocation(location)
	if err != nil {
		return err
	}
	pokemon := extractPokemon(locationInfo)
	fmt.Printf("Exploring %s...\n", *location)
	fmt.Println("Found Pokemon:")
	for _, name := range pokemon {
		fmt.Printf(" - %s\n", name)
	}
	fmt.Println()
	return nil
}

func extractPokemon(locationInfo pokeapi.RespLocationInfo) []string {
	pokemon := make([]string, len(locationInfo.PokemonEncounters))
	for i, encounter := range locationInfo.PokemonEncounters {
		pokemon[i] = encounter.Pokemon.Name
	}
	return pokemon
}
