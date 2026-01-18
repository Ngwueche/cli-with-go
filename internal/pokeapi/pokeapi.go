package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/Ngwueche/cli-with-go.git/pokacache"
)

// pokeBaseURL is the root of the PokeAPI endpoints.
const pokeBaseURL = "https://pokeapi.co/api/v2/"

// Client wraps an http.Client and a cache for API responses.
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient returns a Client configured with a cache and HTTP timeout.
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			// Timeout applies to the entire request, including reads.
			Timeout: time.Minute,
		},
	}
}

