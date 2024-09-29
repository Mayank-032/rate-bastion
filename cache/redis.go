package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	ctx    context.Context
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
	redisObj := r.Client.Get(r.ctx, "")
	if redisObj == nil {
		return "", errors.New("invalid key")
	}

	if redisObj.Err() != nil {
		return "", redisObj.Err()
	}

	return redisObj.Val(), nil
}

func (r *redisCache) Set(key string, value interface{}) error {
	redisObj := r.Client.Set(r.ctx, key, value, 0)
	if redisObj == nil {
		return errors.New("invalid key")
	}

	if redisObj.Err() != nil {
		return redisObj.Err()
	}

	return nil
}

func (r *redisCache) Delete(key string) error {
	redisObj := r.Client.Del(r.ctx, key)
	if redisObj == nil {
		return errors.New("invalid key")
	}

	if redisObj.Err() != nil {
		return redisObj.Err()
	}

	return nil
}
