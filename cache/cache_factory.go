package cache

import (
	"log"
	"strconv"

	"github.com/Mayank-032/rateBastion/enums"
)

var CacheInstance Cache

func NewCache(cacheType enums.CacheType, host, port, database string) (Cache, error) {
	switch cacheType {
	case 1:
		db, err := strconv.Atoi(database)
		if err != nil {
			log.Println("error while string to int conversion: ", err.Error())
			return nil, err
		}

		CacheInstance, err = newRedisCache(host, port, db)
		if err != nil {
			return nil, err
		}
	case 2:
		var err error
		CacheInstance, err = newMemCache(host, port, database)
		if err != nil {
			return nil, err
		}
	default:
		return nil, nil
	}

	return CacheInstance, nil
}
