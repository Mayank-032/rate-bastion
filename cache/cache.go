package cache

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Delete(key string) error
}
