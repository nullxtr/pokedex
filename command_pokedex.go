package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		return errors.New("pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
