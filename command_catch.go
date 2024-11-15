package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	pokemon := args[0]
	if len(args) != 1 {
		return errors.New("give Pokemon name")
	}
	// get Pokemon Stats
	pokemonStatsResp, err := cfg.pokeapiClient.PokemonStats(pokemon)
	if err != nil {
		return err
	}

	baseExp := pokemonStatsResp.BaseExperience

	if baseExp == 0 {
		fmt.Println("There is no Pokemon with this Name")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	if getOutcome(float64(baseExp)) {
		fmt.Printf("\033[32m%v was caught!\n\033[0m", pokemon)
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}
	cfg.caughtPokemon[pokemonStatsResp.Name] = pokemonStatsResp
	return nil
}

// fmt.Println("\033[32mThis is green text\033[0m")
func probability(baseStat float64) float64 {
	return (-3.0/1400.0)*baseStat + 1.0357
}

func getOutcome(baseStat float64) bool {
	prob := probability(baseStat)
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Float64()

	return randomValue <= prob
}
