package ratelimiter

import "time"

type slidingWindow struct {
	MaxRequestsInTimeWindow int
	TimeWindowInSeconds     int
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

	newRequestTimestamp := time.Now()                            // current time
	initialRequestTimestamp := userRequestsTimestamps[userId][0] // first request of the user
	timeDiff := newRequestTimestamp.Sub(initialRequestTimestamp).Seconds()

	// if the time difference is greater than the time window, remove all the initial request timestamps whose time difference is greater than the time window
	if timeDiff > float64(s.TimeWindowInSeconds) {
		for len(userRequestsTimestamps[userId]) > 0 && timeDiff > float64(s.TimeWindowInSeconds) {
			initialRequestTimestamp = userRequestsTimestamps[userId][0]
			timeDiff = newRequestTimestamp.Sub(initialRequestTimestamp).Seconds()
			if timeDiff > float64(s.TimeWindowInSeconds) {
				userRequestsTimestamps[userId] = userRequestsTimestamps[userId][1:]
			}
		}
	} else {
		// if the time difference is less than the time window, check if the number of requests is greater than the max requests in the time window
		if len(userRequestsTimestamps[userId]) >= s.MaxRequestsInTimeWindow {
			return false
		}
	}

	userRequestsTimestamps[userId] = append(userRequestsTimestamps[userId], newRequestTimestamp)
	return true
}
