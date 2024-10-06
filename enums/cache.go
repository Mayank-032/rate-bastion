package enums

type CacheType int

const (
	Memory CacheType = iota
	REDIS
	MEMCACHE
)
