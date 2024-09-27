package ratelimiter

func NewRateLimiter(strategy string, maxRequestsInTimeWindow, timeWindowInSeconds int) RateLimiter {
	switch strategy {
	case "token_bucket":
		return newTokenBucketRateLimiter(maxRequestsInTimeWindow, timeWindowInSeconds)
	case "sliding_window":
		return newSlidingWindowRateLimiter(maxRequestsInTimeWindow, timeWindowInSeconds)
	default:
		return nil
	}
}
