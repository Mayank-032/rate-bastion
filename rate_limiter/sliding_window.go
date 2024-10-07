package ratelimiter

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Mayank-032/rate-bastion/cache"
)

type slidingWindow struct {
	MaxRequestsInTimeWindow int
	TimeWindowInSeconds     int
}

func newSlidingWindowRateLimiter(maxRequestsInTimeWindow int, timeWindowInSeconds int) RateLimiter {
	return &slidingWindow{
		MaxRequestsInTimeWindow: maxRequestsInTimeWindow,
		TimeWindowInSeconds:     timeWindowInSeconds,
	}
}

type userTimestamps struct {
	Timestamps []time.Time `json:"timestamps"`
}

func (s *slidingWindow) IsRequestAllowed(userId string) (bool, error) {
	var cacheInstance = cache.CacheInstance
	var user = userTimestamps{}

	userObj, err := cacheInstance.Get(userId)
	if err != nil {
		log.Printf("err: %v; unable to get user instance from cache\n", err.Error())

		if !strings.EqualFold(err.Error(), "invalid key") {
			return false, err
		}

		user = userTimestamps{Timestamps: make([]time.Time, 0)}
	} else {
		err = json.Unmarshal([]byte(userObj), &user)
		if err != nil {
			log.Printf("err: %v, unable to unmarshal bytes to user struct\n", err.Error())
			return false, err
		}
	}

	var newRequestTimestamp = time.Now() // current time
	var requests = user.Timestamps       // first request of the user

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	// remove outdated request timestamps
	var cutOffTime = newRequestTimestamp.Add(-time.Duration(s.TimeWindowInSeconds) * time.Second)
	for len(requests) > 0 && requests[0].Before(cutOffTime) { // timestamps before cutoff time are outdated
		requests = requests[1:]
	}

	// if the number of requests is greater than the max requests in the time window, return false
	if len(requests) >= s.MaxRequestsInTimeWindow {
		userBytes, err := json.Marshal(user)
		if err != nil {
			log.Printf("err: %v, unable to marshal struct\n", err.Error())
			return false, err
		}
		err = cacheInstance.Set(userId, string(userBytes))
		if err != nil {
			log.Printf("err: %v, unable to set key-value pair", err.Error())
			return false, err
		}
		return false, nil
	}

	// add the new request timestamp
	user.Timestamps = append(requests, newRequestTimestamp)
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("err: %v, unable to marshal struct\n", err.Error())
		return false, err
	}
	err = cacheInstance.Set(userId, string(userBytes))
	if err != nil {
		log.Printf("err: %v, unable to set key-value pair", err.Error())
		return false, err
	}

	return true, nil
}
