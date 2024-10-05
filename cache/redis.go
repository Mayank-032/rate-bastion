package cache

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	ctx    context.Context
	Client *redis.Client
}

func newRedisCache(host, port string, db int) (Cache, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   db,
	})

	res, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	log.Printf("PING-%v\n", res)

	return &redisCache{
		ctx:    ctx,
		Client: client,
	}, nil
}

func (r *redisCache) Get(key string) (string, error) {
	if r.Client == nil {
		return "", errors.New("redis client is not initialized")
	}

	value, err := r.Client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("invalid key")
		}
		if err == redis.ErrClosed {
			return "", errors.New("connection closed")
		}
		return "", err
	}

	return value, nil
}

func (r *redisCache) Set(key string, value interface{}) error {
	if r.Client == nil {
		return errors.New("redis client is not initialized")
	}

	res, err := r.Client.Set(r.ctx, key, value, 0).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("invalid key")
		}
		if err == redis.ErrClosed {
			return errors.New("connection closed")
		}

		return err
	}
	log.Printf("user: %v, res: %v", key, res)

	return nil
}

func (r *redisCache) Delete(key string) error {
	if r.Client == nil {
		return errors.New("redis client is not initialized")
	}

	res, err := r.Client.Del(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("invalid key")
		}
		return err
	}
	log.Println("res: ", res)

	return nil
}
