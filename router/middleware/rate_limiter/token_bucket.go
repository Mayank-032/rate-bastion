package ratelimiter

import (
	"sync"
	"time"
)

type tokenBucket struct {
	MaxRequestsInTimeWindow int       // max number of requests in the time window (It also refers to number of token/requests allowed in every TimeWindowInSeconds)
	TimeWindowInSeconds     int       // time after which the token bucket is refilled
	TokensInBucket          int       // current number of tokens in the bucket
	LastRefillTime          time.Time // bucket, last refilled at
	mu                      sync.Mutex
}

var userBuckets = make(map[string]*tokenBucket, 0)

func NewTokenBucketRateLimiter(maxRequestsInTimeWindow int, timeWindowInSeconds int) RateLimiter {
	return &tokenBucket{
		MaxRequestsInTimeWindow: maxRequestsInTimeWindow,
		TimeWindowInSeconds:     timeWindowInSeconds,
		LastRefillTime:          time.Now(),
		TokensInBucket:          0,
	}
}

func (t *tokenBucket) IsRequestAllowed(userId string) bool {
	// check if the user has a bucket, if not create one
	if _, ok := userBuckets[userId]; !ok {
		userBuckets[userId] = t
	}
	bucket := userBuckets[userId]

	// get current time
	currentTimestamp := time.Now()

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	// get time since last refill
	timeSinceLastRefill := currentTimestamp.Sub(bucket.LastRefillTime).Seconds()

	// calculate if it passed the time window, recharge the token bucket with number of tokens not exceeding the MaxRequestsInTimeWindow
	if timeSinceLastRefill >= float64(bucket.TimeWindowInSeconds) {
		bucket.TokensInBucket = bucket.MaxRequestsInTimeWindow
		bucket.LastRefillTime = currentTimestamp
	}

	// check if there are tokens in the bucket
	if bucket.TokensInBucket <= 0 {
		return false
	}

	// decrement the number of tokens in the bucket
	bucket.TokensInBucket--

	return true
}
