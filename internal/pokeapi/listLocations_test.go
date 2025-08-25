package pokeapi

import (
	"fmt"
	"testing"
	"time"

	"github.com/Awowz/Pokedex/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}
	for index, data := range cases {
		t.Run(fmt.Sprintf("Test case %v", index), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(data.key, data.val)
			val, ok := cache.Get(data.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(data.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
