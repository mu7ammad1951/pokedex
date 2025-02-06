package main

import "fmt"

func commandInspect(cfg *config, arg *string) error {
	if arg == nil {
		return fmt.Errorf("error: pokemon argument missing 'inspect <pokemon-name>'")
	}

	pokemonInfo, exists := cfg.Pokedex[*arg]
	if !exists {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonInfo.Types {
		fmt.Printf(" - %v\n", pokeType.Type.Name)
	}
	return nil
}
