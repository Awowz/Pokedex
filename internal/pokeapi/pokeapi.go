package pokeapi

import (
	"fmt"
	"math/rand"
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

type PokemonData struct {
	Base_experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Stats           []struct {
		Base_stat int `json:"base_stat"`
		Stat      struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
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

func (p PokemonData) AttemptCatch() bool {
	catchrate := CATCH_RATE - p.Base_experience
	if catchrate <= 0 {
		catchrate = 1
	}
	caughtChance := rand.Intn(CATCH_RATE)
	return caughtChance <= catchrate
}
