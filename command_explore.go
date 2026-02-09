package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	name := args[0]
	locationExploreResp, err := cfg.pokeapiClient.LocationExplore(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	for _, pokemonEncounters := range locationExploreResp.PokemonEncounters {
		fmt.Println(pokemonEncounters.Pokemon.Name)
	}
	return nil
}