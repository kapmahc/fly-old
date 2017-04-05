package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// Handle handle
type Handle func(*Context) error

// NewRouter new router
func NewRouter(opt render.Options) *Router {
	return &Router{
		root:    mux.NewRouter(),
		render:  render.New(opt),
		handles: make([]Handle, 0),
	}
}

// Router router
type Router struct {
	root    *mux.Router
	render  *render.Render
	handles []Handle
}

// Use use
func (p *Router) Use(handles ...Handle) {
	p.handles = append(p.handles, handles...)
}

// LoadHTML load html
func (p *Router) LoadHTML(string) {
	// TODO
}

// Walk walk routes
func (p *Router) Walk(func(name, method, path string)) {
	// TODO
}

// Group group
func (p *Router) Group(path string, fn func(*Router), handles ...Handle) {
	rt := p.root.PathPrefix(path).Subrouter()
	items := make([]Handle, len(p.handles))
	copy(items, p.handles)
	items = append(items, handles...)

	fn(&Router{
		root:    rt,
		render:  p.render,
		handles: items,
	})
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
	items := make([]Handle, len(p.handles))
	copy(items, p.handles)
	items = append(items, handles...)

	p.root.HandleFunc(path, func(wrt http.ResponseWriter, req *http.Request) {
		ctx := Context{Request: req, Writer: wrt}
		for _, hnd := range items {
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
