package fly_test

import (
	"testing"

	"github.com/kapmahc/fly"
)

func TestHmac(t *testing.T) {
	en := fly.Hmac{Key: []byte("123456")}
	hello := "Hello, Champak!"
	code := en.Sum([]byte(hello))
	t.Log(string(code))
	if !en.Chk([]byte(hello), code) {
		t.Fatalf("check password failed")
	}

}
