package loader

import (
	"context"
	"log"

	"github.com/dgraph-io/ristretto"
	"github.com/graph-gophers/dataloader"
)

// ristrettoCache implements the dataloader.Cache interface
type ristrettoCache struct{ ristretto *ristretto.Cache }

func (c *ristrettoCache) Get(_ context.Context, key dataloader.Key) (dataloader.Thunk, bool) {
	v, ok := c.ristretto.Get(key.String())
	if ok {
		return v.(dataloader.Thunk), ok
	}
	return nil, ok
}

func (c *ristrettoCache) Set(_ context.Context, key dataloader.Key, value dataloader.Thunk) {
	c.ristretto.SetWithTTL(key.String(), value, 1, 3e+11) // five minutes
}

func (c *ristrettoCache) Delete(_ context.Context, key dataloader.Key) bool {
	_, ok := c.ristretto.Get(key.String())
	if ok {
		c.ristretto.Del(key.String())
		return true
	}
	return false
}

func (c *ristrettoCache) Clear() {
	c.ristretto.Clear()
}

// TODO: replace with redis
func newRistrettoCache() dataloader.Cache {
	ristretto, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7, // number of keys to track frequency of
		MaxCost:     1e6, // maximum cost of cache
		BufferItems: 64,  // number of keys per Get buffer
	})
	if err != nil {
		log.Fatalln("Unable to create cache:", err)
	}
	return &ristrettoCache{ristretto}
}
