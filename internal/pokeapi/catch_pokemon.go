package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

// CatchPokemon -
func (c *Client) CatchPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		log.Printf("Getting data from cache for url: %s\n", url)
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemonResp, nil
	}
	log.Printf("Fetching data from API for url: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}
	

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}
	c.cache.Add(url, dat)

	return pokemonResp, nil
}
