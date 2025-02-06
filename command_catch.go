package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg *string) error {
	pokemonInfo, err := cfg.PokeAPIClient.GetPokemonInfo(arg)
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonInfo.Name)

	chance := rand.Intn(700)
	caught := chance > pokemonInfo.BaseExperience

	if caught {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		cfg.Pokedex[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}

	return nil
}
