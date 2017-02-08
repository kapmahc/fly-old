package base

import (
	"time"

	"github.com/astaxie/beego"
	_cache "github.com/astaxie/beego/cache"
)

var cache _cache.Cache

// CachePut put into cache
func CachePut(key string, val interface{}, ttl time.Duration) {
	if err := cache.Put(key, val, ttl); err != nil {
		beego.Error(err)
	}
}

// CacheGet get from cache
func CacheGet(key string, val interface{}, ttl time.Duration) interface{} {
	return cache.Get(key)
}

func init() {
	var err error
	cache, err = _cache.NewCache("redis", beego.AppConfig.String("cacheprovider"))
	if err != nil {
		beego.Error(err)
	}
}
