package fly

import (
	"bytes"
	"encoding/gob"
)

// Marshal object to binrary
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

// Unmarshal binrary to object
func Unmarshal(b []byte, v interface{}) error {
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(b)
	return dec.Decode(v)
}
