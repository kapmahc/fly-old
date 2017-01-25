package fly_test

import (
	"crypto/aes"
	"testing"

	"github.com/kapmahc/fly"
)

func TestAes(t *testing.T) {
	cip, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	a := fly.Aes{Cip: cip}

	hello := "Hello, Champak!"
	code, err := a.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(code))
	plain, err := a.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("want %s get %s", hello, string(plain))
	}

}
