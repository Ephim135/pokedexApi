 package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	value, exists := cfg.pokeapiCache.Get(cfg.nextLocationsURL)
	if exists {
		// get from cache
		fmt.Print("from Cache")
		locationsResp := value
	} else {
		fmt.Print
		locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
	// add to cache 	
		cfg.pokeapiCache.Add(cfg.nextLocationsURL, locationsResp)
	}
	// set New Url
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	printLocations(locationsResp)

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	value, exists := cfg.pokeapiCache.Get(cfg.prevLocationsURL)
	if exists {
		locationsResp := value
	} else {
		locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		cfg.pokeapiCache.Add(cfg.prevLocationsURL, locationsResp)
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	printLocations(locationResp)

	return nil
}

func printLocations(locationsResp RespShallowLocations) {
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
}