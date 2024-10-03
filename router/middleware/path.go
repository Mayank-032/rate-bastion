package middleware

import (
	"net/http"
	ratelimiter "rateBastion/router/middleware/rate_limiter"
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
		userId := r.Header.Get("X-User-Id")
		if userId == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// get rate limiter instance from factory
		rateLimiter := ratelimiter.NewRateLimiter("token_bucket", 2, 10)

		// check if user id is in the rate limiter and if the request is allowed
		ok, err := rateLimiter.IsRequestAllowed(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !ok {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// if request is allowed, serve the request
		next.ServeHTTP(w, r)
	})
}
