package redis_test

import (
	"testing"

	"github.com/kapmahc/fly"
	"github.com/kapmahc/fly/redis"
)

func TestQueue(t *testing.T) {
	var que fly.Queue
	que = &redis.Queue{Redis: &pool, Namespace: namespace}
	if err := que.Send("echo", []byte("hello, queue")); err != nil {
		t.Fatal(err)
	}
	name, buf, err := que.Receive("echo", "ping")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("get task from %s: %s", name, string(buf))
}
