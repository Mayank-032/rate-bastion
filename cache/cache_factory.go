package cache

import (
	"log"
	"rate-limiter/configs"
	"strconv"
)

func NewCache(cacheType string) Cache {
	switch cacheType {
	case "redis":
		db, err := strconv.Atoi(configs.Configuration.Redis.Database)
		if err != nil {
			log.Println("error while string to int conversion: ", err.Error())
			return nil
		}
		return newRedisCache(configs.Configuration.Redis.Host, configs.Configuration.Redis.Port, db)
	case "memcache":
		return newMemCache(configs.Configuration.Redis.Host, configs.Configuration.Redis.Port, configs.Configuration.Redis.Database)
	default:
		return nil
	}
}
