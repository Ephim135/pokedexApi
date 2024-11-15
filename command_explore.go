package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("to many arguments")
	}
	if len(args) < 1 {
		return errors.New("to few arguments")
	}

	pokemonResp, err := cfg.pokeapiClient.ListPokemons(args[0])
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
