package redis

import (
	"fmt"

	_redis "github.com/garyburd/redigo/redis"
)

// Queue queue
type Queue struct {
	Redis     *_redis.Pool `inject:""`
	Namespace string       `inject:"namespace"`
}

// Send send a task message
func (p *Queue) Send(name string, task []byte) error {
	c := p.Redis.Get()
	defer c.Close()
	_, err := c.Do("LPUSH", p.key(name), task)
	return err
}

// Receive receive a task
func (p *Queue) Receive(args ...string) (string, []byte, error) {
	c := p.Redis.Get()
	defer c.Close()
	var names []interface{}
	for _, v := range args {
		names = append(names, p.key(v))
	}
	names = append(names, 0)
	val, err := _redis.ByteSlices(c.Do("BRPOP", names...))
	if err != nil {
		return "", nil, err
	}
	return string(val[0])[len(p.key("")):], val[1], nil
}

func (p *Queue) key(k string) string {
	return fmt.Sprintf("%s@queue://%s", p.Namespace, k)
}
