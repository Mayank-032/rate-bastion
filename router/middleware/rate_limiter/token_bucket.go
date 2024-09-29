package ratelimiter

import (
	"encoding/json"
	"log"
	"rate-limiter/cache"
	"strings"
	"sync"
	"time"
)

type tokenBucket struct {
	MaxRequestsInTimeWindow int // max number of requests in the time window (It also refers to number of token/requests allowed in every TimeWindowInSeconds)
	TimeWindowInSeconds     int // time after which the token bucket is refilled
}

func newTokenBucketRateLimiter(maxRequestsInTimeWindow int, timeWindowInSeconds int) RateLimiter {
	return &tokenBucket{
		MaxRequestsInTimeWindow: maxRequestsInTimeWindow,
		TimeWindowInSeconds:     timeWindowInSeconds,
	}
}

type userBucket struct {
	TokensInBucket int       `json:"tokens_in_bucket"`
	LastRefillTime time.Time `json:"last_refill_time"`
}

func (t *tokenBucket) IsRequestAllowed(userId string) (bool, error) {
	// check if the user has a bucket, if not create one
	var cacheInstance = cache.CacheInstance
	var user = userBucket{}

	userObj, err := cacheInstance.Get(userId)
	if err != nil {
		log.Printf("err: %v; unable to get user instance from cache\n", err.Error())

		if !strings.EqualFold(err.Error(), "invalid key") {
			return false, err
		}

		user = userBucket{TokensInBucket: 0, LastRefillTime: time.Now()}
	} else {
		err = json.Unmarshal([]byte(userObj), &user)
		if err != nil {
			log.Printf("err: %v, unable to unmarshal bytes to user struct\n", err.Error())
			return false, err
		}
	}

	// get current time
	currentTimestamp := time.Now()

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	// get time since last refill
	timeSinceLastRefill := currentTimestamp.Sub(user.LastRefillTime).Seconds()

	// calculate if it passed the time window, recharge the token bucket with number of tokens not exceeding the MaxRequestsInTimeWindow
	if timeSinceLastRefill >= float64(t.TimeWindowInSeconds) {
		user.TokensInBucket = t.MaxRequestsInTimeWindow
		user.LastRefillTime = currentTimestamp
	}

	// check if there are tokens in the bucket
	if user.TokensInBucket <= 0 {
		userBytes, err := json.Marshal(user)
		if err != nil {
			log.Printf("err: %v, unable to marshal struct\n", err.Error())
			return false, err
		}
		cacheInstance.Set(userId, string(userBytes))
		return false, nil
	}

	// decrement the number of tokens in the bucket
	user.TokensInBucket--

	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("err: %v, unable to marshal struct\n", err.Error())
		return false, err
	}
	cacheInstance.Set(userId, string(userBytes))

	return true, nil
}
