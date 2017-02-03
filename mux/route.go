package mux

import "net/http"

type route struct {
	name     string
	pattern  string
	method   string
	handlers []interface{}
}

func (p *route) match(r *http.Request) bool {
	if r.Method != p.method {
		return false
	}
	return true
}
