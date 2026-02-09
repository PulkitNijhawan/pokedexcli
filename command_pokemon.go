package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	name := args[0]
	if _, ok := cfg.Pokedex[name]; ok {
		fmt.Printf("%s already caught! \n", name)
		return nil
	}
	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	const catchPower = 50
    res := rand.Intn(pokemonResp.BaseExperience)

    if res < catchPower {
		cfg.Pokedex[name] = pokemonResp
        fmt.Printf("%s was caught!\n", name)
    } else {
        fmt.Printf("%s escaped!\n", name)
    }
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	name := args[0]
	if p, ok := cfg.Pokedex[name]; ok {
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Height: %d\n", p.Height)
		fmt.Printf("Weight: %d\n", p.Weight)

		fmt.Println("Stats:")
		for _, s := range p.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.Pokedex {
		fmt.Printf("  - %s\n", k)
	}
	
	return nil
}