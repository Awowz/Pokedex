package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ShallowListLocations(providedURL *string) (PokeMap, error) {
	url := BASE_URL + URL_LOCATIONS
	if providedURL != nil {
		url = *providedURL
	}

	myCache, valid := c.clientCache.Get(url)
	if valid {
		myPokeMap, err := GetPokeMapFromCache(myCache)
		if err != nil {
			return PokeMap{}, err
		}
		return myPokeMap, nil
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
	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeMap{}, fmt.Errorf("io could not read response: %s", err)
	}

	var pokeMapData PokeMap
	err = json.Unmarshal(bodyByte, &pokeMapData)
	if err != nil {
		return PokeMap{}, fmt.Errorf("could not decode pokemonAPI json. Error: %v", err)
	}

	c.clientCache.Add(url, bodyByte)
	return pokeMapData, nil
}

func GetPokeMapFromCache(data []byte) (PokeMap, error) {
	var myPokeMap PokeMap
	err := json.Unmarshal(data, &myPokeMap)
	if err != nil {
		return PokeMap{}, err
	}
	return myPokeMap, nil
}

func (c *Client) ListPokemonFromLocation(providedURL string) (PokeMapEncounters, error) {

	myCache, valid := c.clientCache.Get(providedURL)
	if valid {
		myPokeMap, err := GetPokeMapPokemonFromCache(myCache)
		if err != nil {
			return PokeMapEncounters{}, err
		}
		return myPokeMap, nil
	}

	req, err := http.NewRequest("GET", providedURL, nil)
	if err != nil {
		return PokeMapEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeMapEncounters{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokeMapEncounters{}, fmt.Errorf("requested page not found")
	}

	bodyBte, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeMapEncounters{}, err
	}

	var pokeEncounters PokeMapEncounters
	err = json.Unmarshal(bodyBte, &pokeEncounters)
	if err != nil {
		return PokeMapEncounters{}, fmt.Errorf("error unmarshaling json data, %s", err)
	}

	c.clientCache.Add(providedURL, bodyBte)
	return pokeEncounters, nil
}

func GetPokeMapPokemonFromCache(data []byte) (PokeMapEncounters, error) {
	var myPokeMap PokeMapEncounters
	err := json.Unmarshal(data, &myPokeMap)
	if err != nil {
		return PokeMapEncounters{}, err
	}
	return myPokeMap, nil
}
