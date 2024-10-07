package rateBastion

import (
	"errors"
	"log"

	"github.com/Mayank-032/rate-bastion/cache"
	"github.com/Mayank-032/rate-bastion/configs"
	rateLimiterStrategy "github.com/Mayank-032/rate-bastion/limiting_strategy"
)

func NewRateLimiter(config *configs.Config) (rateLimiterStrategy.LimitingMethod, error) {
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
	var rateLimiterInstance rateLimiterStrategy.LimitingMethod
	switch config.Strategy {
	case 1: // TOKEN_BUCKET
		rateLimiterInstance = rateLimiterStrategy.NewTokenBucketRateLimiter(config.MaxRequestsAllowedInTimeWindow, config.TimeWindowInSeconds)
	case 2: // SLIDING_WINDOW_LOG
		rateLimiterInstance = rateLimiterStrategy.NewSlidingWindowRateLimiter(config.MaxRequestsAllowedInTimeWindow, config.TimeWindowInSeconds)
	default:
		return nil, errors.New("unknown rate limiter strategy")
	}

	return rateLimiterInstance, nil
}
