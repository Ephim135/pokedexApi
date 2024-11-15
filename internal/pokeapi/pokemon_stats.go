package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonStats(pokemon string) (RespShallowPokemonStats, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemonStats{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemonStats{}, nil
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonStats{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonStats{}, err
	}

	pokemonResp := RespShallowPokemonStats{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemonStats{}, nil
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil

}
