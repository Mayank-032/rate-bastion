package cache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type memCache struct {
	Client *memcache.Client
}

func newMemCache(host, port, _ string) Cache {
	servers := []string{fmt.Sprintf("%v:%v", host, port)}
	return &memCache{
		Client: memcache.New(servers...),
	}
}

func (m *memCache) Get(key string) (string, error) {
	return "", nil
}

func (m *memCache) Set(key string, value string) error {
	return nil
}

func (m *memCache) Delete(key string) error {
	return nil
}