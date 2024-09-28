package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	Client *redis.Client
}

func newRedisCache(host, port string, db int) Cache {
	return &redisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
			DB:   db,
		}),
	}
}

func (r *redisCache) Get(key string) (string, error) {
	return "", nil
}

func (r *redisCache) Set(key, value string) error {
	return nil
}

func (r *redisCache) Delete(key string) error {
	return nil
}
