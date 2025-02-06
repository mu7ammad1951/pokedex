package main

import (
	"time"

	"github.com/mu7ammad1951/pokedex/internal/pokeapi"
)

func main() {
	cfg := config{
		PokeAPIClient: pokeapi.NewClient(5*time.Second, 20*time.Second),
		Pokedex:       map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
