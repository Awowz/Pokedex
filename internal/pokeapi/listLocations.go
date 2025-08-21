package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ShallowListLocations(providedURL *string) (PokeMap, error) {
	url := BASE_URL + URL_LOCATIONS
	if providedURL != nil {
		url = *providedURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeMap{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeMap{}, err
	}
	defer res.Body.Close()

	var pokeMapData PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeMapData)

	if err != nil {
		return PokeMap{}, fmt.Errorf("could not decode pokemonAPI json. Error: %v", err)
	}

	return pokeMapData, nil
}
