package ratelimiter

func NewRateLimiter(strategy string, maxRequestsInTimeWindow, timeWindowInSeconds int) RateLimiter {
	switch strategy {
	case "token_bucket":
		return NewTokenBucketRateLimiter(maxRequestsInTimeWindow, timeWindowInSeconds)
	case "sliding_window":
		return NewSlidingWindowRateLimiter(maxRequestsInTimeWindow, timeWindowInSeconds)
	default:
		return nil
	}
}