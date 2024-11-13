package main

import (
	"time"

	"github.com/Ephim135/pokedexApi/internal/pokeapi"
	"github.com/Ephim135/pokedexApi/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokeapiCache: pokeCache
	}


	startRepl(cfg)
}
