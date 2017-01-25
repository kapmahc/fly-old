package redis_test

import (
	"testing"
	"time"

	_redis "github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly"
	"github.com/kapmahc/fly/redis"
)

type S struct {
	IV int
	SV string
}

var pool = _redis.Pool{
	MaxIdle:     3,
	IdleTimeout: 240 * time.Second,
	Dial: func() (_redis.Conn, error) {
		return _redis.Dial(
			"tcp",
			"localhost:6379",
		)
	},
}

var namespace = "test"

func TestCache(t *testing.T) {
	var c fly.Cache
	c = &redis.Cache{Redis: &pool, Namespace: namespace, Coder: &fly.GobCoder{}}

	s1 := S{IV: 100, SV: "hello, champak!"}
	if err := c.Set("hello", &s1, 60*time.Minute); err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := c.Get("hello", &s2); err == nil {
		if s1.IV != s2.IV || s1.SV != s2.SV {
			t.Fatalf("want %v get %v", s1, s2)
		}
	} else {
		t.Fatal(err)
	}
	if keys, err := c.Keys(); err == nil {
		t.Logf("keys: %v", keys)
	} else {
		t.Fatal(err)
	}

	c.Set("aaa", 111, time.Hour*2)
	if err := c.Delete("aaa"); err != nil {
		t.Fatal(err)
	}

	if err := c.Clear(); err != nil {
		t.Fatal(err)
	}
}
