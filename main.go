package main

import (
	"github.com/nullxtr/pokedexcli/interntal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemons: map[string]pokeapi.Pokemon{},
		pokeapiClient:  pokeClient,
	}

	startRepl(cfg)
}
