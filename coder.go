package fly

import (
	"bytes"
	"encoding/gob"
)

// Coder coder
type Coder interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte, v interface{}) error
}

// GobCoder coder by gob
type GobCoder struct {
}

// Marshal object to binrary
func (p *GobCoder) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

// Unmarshal binrary to object
func (p *GobCoder) Unmarshal(b []byte, v interface{}) error {
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(b)
	return dec.Decode(v)
}
