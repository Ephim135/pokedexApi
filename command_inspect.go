package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("give a Pokemon-Name")
	}

	pokemon := args[0]
	val, ok := cfg.caughtPokemon[pokemon]
	if ok {
		fmt.Printf("Name: %v\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Printf("STATS:\n")
		for _, element := range val.Stats {
			fmt.Printf("   -%v: %v\n", element.Stat.Name, element.BaseStat)
		}
		fmt.Printf("TYPES:\n")
		for _, element := range val.Types {
			fmt.Printf("   - %v\n", element.Type.Name)
		}
	} else {
		fmt.Println("pokemon not caught yet")
	}
	return nil
}
