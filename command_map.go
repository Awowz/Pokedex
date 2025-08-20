package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(conf *config) error {
	var res *http.Response
	var err error
	if conf.next != nil {
		res, err = http.Get(*conf.next)
	} else {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
	}
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var pokeMapData PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeMapData)

	if err != nil {
		return fmt.Errorf("could not decode pokemonAPI json. Error: %v", err)
	}

	conf.next = pokeMapData.Next
	conf.previous = pokeMapData.Previous

	for _, x := range pokeMapData.Results {
		fmt.Println(x.Name)
	}

	return nil
}
