package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handle handle
type Handle func(*Context) error

// Router router
type Router struct {
	root *mux.Router
}

// Group group
func (p *Router) Group(path string, fn func(*Router), handles ...Handle) {
	rt := p.root.PathPrefix(path).Subrouter()
	fn(&Router{root: rt})
}

// Get http aget
func (p *Router) Get(name, path string, handles ...Handle) {
	p.add(http.MethodGet, name, path, handles...)
}

// Post http post
func (p *Router) Post(name, path string, handles ...Handle) {
	p.add(http.MethodPost, name, path, handles...)
}

// Delete http delete
func (p *Router) Delete(name, path string, handles ...Handle) {
	p.add(http.MethodDelete, name, path, handles...)
}

func (p *Router) add(method, name, path string, handles ...Handle) {
	p.root.HandleFunc(path, func(wrt http.ResponseWriter, req *http.Request) {
		ctx := Context{Request: req, Writer: wrt}
		for _, hnd := range handles {
			if err := hnd(&ctx); err != nil {
				ctx.Abort(http.StatusInternalServerError, err)
				return
			}
			if ctx.written {
				return
			}
		}
	}).Methods(method)
}
