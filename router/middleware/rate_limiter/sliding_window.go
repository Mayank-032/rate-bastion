package ratelimiter

import (
	"sync"
	"time"
)

type slidingWindow struct {
	MaxRequestsInTimeWindow int
	TimeWindowInSeconds     int
	mu                      sync.Mutex
}

var userRequestsTimestamps = make(map[string][]time.Time, 0)

func NewSlidingWindowRateLimiter(maxRequestsInTimeWindow int, timeWindowInSeconds int) RateLimiter {
	return &slidingWindow{
		MaxRequestsInTimeWindow: maxRequestsInTimeWindow,
		TimeWindowInSeconds:     timeWindowInSeconds,
	}
}

func (s *slidingWindow) IsRequestAllowed(userId string) bool {
	if _, ok := userRequestsTimestamps[userId]; !ok {
		userRequestsTimestamps[userId] = make([]time.Time, 0)
	}

	newRequestTimestamp := time.Now()          // current time
	requests := userRequestsTimestamps[userId] // first request of the user

	// remove outdated request timestamps
	s.mu.Lock()
	defer s.mu.Unlock()

	cutOffTime := newRequestTimestamp.Add(-time.Duration(s.TimeWindowInSeconds) * time.Second)
	for len(requests) > 0 && requests[0].Before(cutOffTime) { // timestamps before cutoff time are outdated
		requests = requests[1:]
	}

	// if the number of requests is greater than the max requests in the time window, return false
	if len(requests) >= s.MaxRequestsInTimeWindow {
		return false
	}

	// add the new request timestamp
	userRequestsTimestamps[userId] = append(requests, newRequestTimestamp)
	return true
}
