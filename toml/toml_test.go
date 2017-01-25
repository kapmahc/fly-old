package toml_test

import (
	"os"
	"testing"
	"time"

	"github.com/kapmahc/fly/toml"
)

func TestConfig(t *testing.T) {
	const name = "test.toml"
	os.Remove(name)
	cfg := map[string]interface{}{
		"I": 123,
		"S": "hello",
		"F": 1.23,
		"T": time.Now(),
		"H": map[string]interface{}{"I": 123},
	}
	if err := toml.Write(name, cfg); err != nil {
		t.Fatal(err)
	}

	tmp, err := toml.Read(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tmp)
	t.Logf("%s %d", tmp["T"].(time.Time), tmp["H"].(map[string]interface{})["I"].(int64))
}
