// Package cache implements a simple thread-safe cache with expiration.
package cache

import (
	"reflect"
	"sync"
	"time"

	"github.com/cespare/xxhash"
)

// Item represents a saved object in the cache. The object has a TTL that the
// clean up routine uses to check if the object has expired.
type Item struct {
	object interface{}
	ttl    time.Time
}

// Cache is a thread-safe cache that holds objects with expiration.
type Cache struct {
	expiration    time.Duration
	items         map[uint64]Item
	mutex         sync.RWMutex
	maxObjectSize uint64
}

// hash wraps the xxHash function.
func hash(key string) uint64 {
	return xxhash.Sum64String(key)
}

// Set adds the object to the cache using the provided key. The key is hashed
// with a hash function before insertion.
func (c *Cache) Set(key string, object interface{}) {
	sum := hash(key)

	c.mutex.Lock()
	// Refuse to fill the cache with objects larger than cache.maxObjectSize.
	if uint64(reflect.TypeOf(object).Size()) <= c.maxObjectSize {
		c.items[sum] = Item{
			object: object,
			ttl:    time.Now().Add(c.expiration),
		}
	}
	c.mutex.Unlock()
}

// Get returns a object from the cache using the provided key. The key is hashed
// using a hash function before lookup.
func (c *Cache) Get(key string) (object interface{}, ok bool) {
	sum := hash(key)

	c.mutex.RLock()
	item, ok := c.items[sum]
	if ok {
		object = item.object
		item.ttl = time.Now().Add(c.expiration)
		c.items[sum] = item
	}
	c.mutex.RUnlock()

	return
}

// clean removes outdated items in the cache.
func (c *Cache) clean() {
	c.mutex.Lock()
	for key, item := range c.items {
		if item.ttl.Before(time.Now()) {
			delete(c.items, key)
		}
	}
	c.mutex.Unlock()
}

// New returns a new cache with a clean up routine.
func New(expiration, cleanupInterval time.Duration) (c *Cache) {
	c = &Cache{
		expiration: expiration,
		// Hardcoded max limit of 2000 rows (could be added as an option
		// when starting the server) and max object size of 50 Kb.
		// This will give us about 2000 * 50 Kb = 100 Mb in-memory max
		// size.
		items:         make(map[uint64]Item, 2000),
		maxObjectSize: 50 * 1000,
	}

	go func() {
		for range time.NewTicker(cleanupInterval).C {
			c.clean()
		}
	}()

	return c
}
