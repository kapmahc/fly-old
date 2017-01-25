package fly

import (
	"encoding/xml"
	"io"
)

// Render render
type Render struct {
}

// XML write xml
func (p *Render) XML(w io.Writer, v interface{}) error {
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
	w.Write([]byte("\n"))
	enc := xml.NewEncoder(w)
	return enc.Encode(v)
}
