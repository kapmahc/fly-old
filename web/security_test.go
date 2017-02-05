package web_test

import (
	"testing"

	"github.com/kapmahc/fly/web"
)

func TestSecurity(t *testing.T) {
	en, err := web.NewAesHmacSecurity([]byte("1234567890123456"), []byte("123456"))
	if err != nil {
		t.Fatal(err)
	}
	hello := "Hello, Champak!"
	code, err := en.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(code))
	plain, err := en.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("wang %s get %s", hello, string(plain))
	}

	code = en.Sum([]byte(hello))
	t.Log(string(code))
	if !en.Chk([]byte(hello), code) {
		t.Fatalf("check password failed")
	}
}
