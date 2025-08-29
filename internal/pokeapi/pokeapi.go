package pokeapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Awowz/Pokedex/internal/pokecache"
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

type PokeMapEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Client struct {
	httpClient  http.Client
	clientCache pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient:  http.Client{Timeout: timeout},
		clientCache: pokecache.NewCache(timeout),
	}
}

func (p PokeMapEncounters) DisplayPokemon() {
	fmt.Println("Found Pokemon:")
	for _, pok := range p.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}
}
