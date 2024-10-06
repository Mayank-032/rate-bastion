package ratelimiter

import (
	"errors"
	"log"
	"rateBastion/cache"
	"rateBastion/configs"
)

func NewRateLimiter(config *configs.Config) (RateLimiter, error) {
	// initialise cache for the rate limiter
	cacheInstance, err := cache.NewCache(config.CacheType, config.CacheStore.Host, config.CacheStore.Port, config.CacheStore.Database)
	if err != nil {
		log.Printf("err: %v, unable to initialise cache store\n", err.Error())
		return nil, errors.New("unable to initialise rate limiter")
	}

	if cacheInstance == nil {
		log.Printf("unknown cache store")
		return nil, errors.New("unknown cache store")
	}

	// initialise rate limiter with instance
	var rateLimiter RateLimiter
	switch config.Strategy {
	case 1: // TOKEN_BUCKET
		rateLimiter = newTokenBucketRateLimiter(config.MaxRequestsAllowedInTimeWindow, config.TimeWindowInSeconds)
	case 2: // SLIDING_WINDOW_LOG
		rateLimiter = newSlidingWindowRateLimiter(config.MaxRequestsAllowedInTimeWindow, config.TimeWindowInSeconds)
	default:
		return nil, errors.New("unknown rate limiter strategy")
	}

	return rateLimiter, nil
}
