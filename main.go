package main

import (
	"fmt"
	"log"
	"net/http"
	"rate-limiter/cache"
	"rate-limiter/configs"
	"rate-limiter/router"
)

func main() {
	err := configs.InitConfigurations()
	if err != nil {
		log.Printf("err: %v, unable to load configs", err.Error())
		return
	}
	log.Println("successfully initialised configs")

	// setup cache
	cacheInstance := cache.NewCache("memcache")
	if cacheInstance == nil {
		log.Println("unable to initialise a cache instance")
		return
	}
	log.Println("successfully initialised a cache instance")

	// setup router (pass cache instance while initialising so that our cache will be used along with setting up the router)
	server := router.InitRouter(cacheInstance)

	port := 8000
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), server); err != nil {
		log.Println("unable to start server on port: ", err.Error())
		return
	}

	log.Println("starting server on port: ", port)
}
