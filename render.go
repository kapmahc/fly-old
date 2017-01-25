package fly

import (
	"encoding/xml"
	"io"
)

// H hash
type H map[string]interface{}

// Render render
type Render struct {
	Pretty bool
}

// XML write xml
func (p *Render) XML(w io.Writer, v interface{}) error {
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
	w.Write([]byte("\n"))
	enc := xml.NewEncoder(w)
	if p.Pretty {
		enc.Indent("", "  ")
	}
	return enc.Encode(v)
}
