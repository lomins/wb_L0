package cache

import (
	"log"
	"sync"

	"github.com/lomins/wildberriesL0/internal/models"
)

type Cache struct {
	m  map[string][]byte
	mu *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		m:  make(map[string][]byte),
		mu: new(sync.RWMutex),
	}
}

func (c *Cache) Add(key string, val []byte) {
	if len(key) == 0 {
		log.Println("Cache.Add error, invalid key: ", key)
		return
	}

	c.WithFullLock(func() {
		_, exists := c.m[key]
		if exists {
			log.Println("This key already exists: ", key)
			return
		}

		c.m[key] = val
	})
}

func (c *Cache) Get(key string) (data []byte, found bool) {
	c.WithRLock(func() {
		data, found = c.m[key]
	})
	return data, found
}

func (c *Cache) AddSet(elements []models.Order) {
	for _, element := range elements {
		c.Add(element.ID, element.Data)
	}
}

func (c *Cache) WithRLock(f func()) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	f()
}

func (c *Cache) WithFullLock(f func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	f()
}
