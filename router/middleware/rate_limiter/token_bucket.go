package ratelimiter

type tokenBucket struct {
	MaxRequestsInTimeWindow int
	TimeWindowInSeconds     int
}

func NewTokenBucketRateLimiter(maxRequestsInTimeWindow int, timeWindowInSeconds int) RateLimiter {
	return &tokenBucket{
		MaxRequestsInTimeWindow: maxRequestsInTimeWindow,
		TimeWindowInSeconds:     timeWindowInSeconds,
	}
}

func (t *tokenBucket) IsRequestAllowed(userId string) bool {
	return true
}
