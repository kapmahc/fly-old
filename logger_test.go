package fly_test

import (
	"testing"
	"time"

	"github.com/kapmahc/fly"
)

func TestLogger(t *testing.T) {
	msg := "Hello"
	var log fly.Logger
	log = &fly.ConsoleLogger{}
	for _, lvl := range []int{fly.DEBUG, fly.INFO, fly.WARN, fly.ERROR} {
		log.Level(lvl)
		log.Debug(msg, 123, time.Now())
		log.Info(msg)
		log.Warn(msg)
		log.Error(msg)
	}
}
