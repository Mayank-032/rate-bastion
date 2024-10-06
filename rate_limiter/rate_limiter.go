package ratelimiter

type RateLimiter interface {
	IsRequestAllowed(userId string) (bool, error)
}