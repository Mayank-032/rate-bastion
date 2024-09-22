package middleware

import (
	"net/http"
	ratelimiter "rate-limiter/router/middleware/rate_limiter"
	"strings"
)

func ParseMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(r.Method, method) {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ParseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetch user id from header
		userId := r.Header.Get("X-User-ID")
		if userId == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// check if user id is in the rate limiter and if the request is allowed
		// get rate limiter instance from factory
		rateLimiter := ratelimiter.NewRateLimiter("sliding_window", 1, 2)
		if !rateLimiter.IsRequestAllowed(userId) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// if request is allowed, serve the request
		next.ServeHTTP(w, r)
	})
}
