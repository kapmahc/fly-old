package fly

import "time"

// Cache cache
type Cache interface {
	Set(key string, val interface{}, ttl time.Duration) error
	Get(key string, val interface{}) error
	Keys() ([]string, error)
	Delete(key string) error
	Clear() error
}
