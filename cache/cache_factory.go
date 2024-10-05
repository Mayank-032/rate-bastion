package cache

import (
	"log"
	"os"
	"rateBastion/configs"
	"strconv"
)

var CacheInstance Cache

func NewCache(cacheType string) Cache {
	switch cacheType {
	case "redis":
		db, err := strconv.Atoi(configs.Configuration.Redis.Database)
		if err != nil {
			log.Println("error while string to int conversion: ", err.Error())
			return nil
		}

		var redisHost string
		var redisPort string
		switch configs.Configuration.Environment {
		case "local":
			redisHost = configs.Configuration.Redis.Host
			redisPort = configs.Configuration.Redis.Port
		default:
			redisHost = os.Getenv("REDIS_HOST")
			redisPort = os.Getenv("REDIS_PORT")
		}

		CacheInstance, err = newRedisCache(redisHost, redisPort, db)
		if err != nil {
			log.Printf("err: %v, unable to init redis instance", err.Error())
			return nil
		}
	case "memcache":
		var memcacheHost string
		var memcachePort string
		switch configs.Configuration.Environment {
		case "local":
			memcacheHost = configs.Configuration.Memcache.Host
			memcachePort = configs.Configuration.Memcache.Port
		default:
			memcacheHost = os.Getenv("MEMCACHE_HOST")
			memcachePort = os.Getenv("MEMCACHE_PORT")
		}

		var err error
		CacheInstance, err = newMemCache(memcacheHost, memcachePort, configs.Configuration.Memcache.Database)
		if err != nil {
			log.Printf("err: %v, unable to init memcache instance", err.Error())
			return nil
		}
	default:
		return nil
	}

	return CacheInstance
}
