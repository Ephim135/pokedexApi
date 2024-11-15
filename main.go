package main

import (
	"time"

	"github.com/Ephim135/pokedexApi/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config{
		caughtPokemon: map[string]pokeapi.RespShallowPokemonStats{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
