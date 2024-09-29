package cache

import (
	"log"
	"rate-limiter/configs"
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

		CacheInstance = newRedisCache(configs.Configuration.Redis.Host, configs.Configuration.Redis.Port, db)
	case "memcache":
		CacheInstance = newMemCache(configs.Configuration.Memcache.Host, configs.Configuration.Memcache.Port, configs.Configuration.Memcache.Database)
	default:
		return nil
	}

	return CacheInstance
}
