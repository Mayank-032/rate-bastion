package router

import (
	"net/http"
	"rateBastion/router/middleware"
)

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	successHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success... ;)"))
	})

	// endpoint with rate limiting middleware
	mux.Handle("/limited", middleware.ParseMethod("GET", middleware.ParseHeader(successHandler)))

	// endpoint without rate limiting middleware
	mux.Handle("/unlimited", middleware.ParseMethod("GET", successHandler))

	return mux
}
