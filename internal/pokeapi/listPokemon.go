package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemonDetails(givenPokemon string) (PokemonData, error) {
	url := BASE_URL + URL_POKEMON + "/" + givenPokemon

	myCache, valid := c.clientCache.Get(url)
	if valid {
		myPokeData, err := GetPokemonDetailsFromCache(myCache)
		if err != nil {
			return PokemonData{}, err
		}
		return myPokeData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return PokemonData{}, fmt.Errorf("no pokemon found with that name")
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	var returnedPokemon PokemonData
	err = json.Unmarshal(byteData, &returnedPokemon)
	if err != nil {
		return PokemonData{}, err
	}

	c.clientCache.Add(url, byteData)
	return returnedPokemon, nil
}

func GetPokemonDetailsFromCache(data []byte) (PokemonData, error) {
	var myPoke PokemonData
	err := json.Unmarshal(data, &myPoke)
	if err != nil {
		return PokemonData{}, err
	}
	return myPoke, nil
}
