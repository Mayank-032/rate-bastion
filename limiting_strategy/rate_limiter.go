package ratelimiter

type LimitingMethod interface {
	IsRequestAllowed(userId string) (bool, error)
}
