package cache

import (
	"encoding/json"
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
	memcacheObj, err := m.Client.Get(key)
	if err != nil {
		return "", err
	}
	return string(memcacheObj.Value), nil
}

func (m *memCache) Set(key string, value interface{}) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = m.Client.CompareAndSwap(&memcache.Item{
		Key:   key,
		Value: valueBytes,
	})
	if err != memcache.ErrNotStored {
		return err
	}

	err = m.Client.Set(&memcache.Item{
		Key:   key,
		Value: valueBytes,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *memCache) Delete(key string) error {
	err := m.Client.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
