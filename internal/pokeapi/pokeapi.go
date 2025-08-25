package pokeapi

import (
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
