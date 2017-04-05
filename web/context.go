package web

import "net/http"

// Context context
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	written bool
}

// Abort abort
func (p *Context) Abort(code int, err error) {
	// TODO render templat by different code
	p.written = true
}
