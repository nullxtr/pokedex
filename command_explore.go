package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}

	area := args[0]

	areaResp, err := cfg.pokeapiClient.ExploreArea(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, enc := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
