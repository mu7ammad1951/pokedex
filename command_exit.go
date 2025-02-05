package main

import (
	"fmt"
	"os"
)

func commandExit(confg *config, arg *string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
