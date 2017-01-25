package toml

import (
	"os"

	_toml "github.com/BurntSushi/toml"
)

// Read read from file
func Read(file string) (map[string]interface{}, error) {
	var cfg map[string]interface{}
	if _, err := _toml.DecodeFile(file, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Write write to file
func Write(file string, cfg map[string]interface{}) error {
	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	enc := _toml.NewEncoder(fd)
	return enc.Encode(cfg)
}
