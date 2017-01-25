package redis

import (
	"fmt"
	"time"

	_redis "github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly"
)

// Cache redis cache
type Cache struct {
	Redis     *_redis.Pool `inject:""`
	Coder     fly.Coder    `inject:""`
	Namespace string       `inject:"namespace"`
}

// Set set
func (p *Cache) Set(key string, val interface{}, ttl time.Duration) error {
	buf, err := p.Coder.Marshal(val)
	if err != nil {
		return err
	}

	c := p.Redis.Get()
	defer c.Close()
	_, err = c.Do("SET", p.key(key), buf, "EX", int(ttl/time.Second))
	return err
}

// Get get
func (p *Cache) Get(key string, val interface{}) error {
	c := p.Redis.Get()
	defer c.Close()
	buf, err := _redis.Bytes(c.Do("GET", p.key(key)))
	if err != nil {
		return err
	}
	return p.Coder.Unmarshal(buf, val)
}

// Keys keys
func (p *Cache) Keys() ([]string, error) {
	c := p.Redis.Get()
	defer c.Close()
	keys, err := _redis.Strings(c.Do("KEYS", p.key("*")))
	if err != nil {
		return nil, err
	}
	var val []string
	for _, k := range keys {
		val = append(val, k[len(p.key("")):])
	}
	return val, nil
}

// Delete delete
func (p *Cache) Delete(key string) error {
	c := p.Redis.Get()
	defer c.Close()
	_, e := c.Do("DEL", p.key(key))
	return e
}

// Clear clear
func (p *Cache) Clear() error {
	c := p.Redis.Get()
	defer c.Close()
	keys, err := _redis.Values(c.Do("KEYS", p.key("*")))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

func (p *Cache) key(k string) string {
	return fmt.Sprintf("%s@cache://%s", p.Namespace, k)
}
