package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

// ExploreLocation -
func (c *Client) LocationExplore(name string) (RespLocationExplore, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		log.Printf("Getting data from cache for url: %s\n", url)
		locationExploreResp := RespLocationExplore{}
		err := json.Unmarshal(val, &locationExploreResp)
		if err != nil {
			return RespLocationExplore{}, err
		}

		return locationExploreResp, nil
	}
	log.Printf("Fetching data from API for url: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationExplore{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationExplore{}, err
	}
	

	locationExploreResp := RespLocationExplore{}
	err = json.Unmarshal(dat, &locationExploreResp)
	if err != nil {
		return RespLocationExplore{}, err
	}
	c.cache.Add(url, dat)

	return locationExploreResp, nil
}
