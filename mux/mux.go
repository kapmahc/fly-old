package mux

import (
	"net/http"

	"github.com/kapmahc/fly/inject"
)

// New new mux
func New() *Mux {
	return &Mux{routes: make([]route, 0)}
}

// Mux mux
type Mux struct {
	inject.Injector
	routes []route
}

func (p *Mux) ServeHTTP(wrt http.ResponseWriter, req *http.Request) {
	for _, rt := range p.routes {
		if rt.match(req) {
			for _, h := range rt.handlers {
				if _, err := p.Invoke(h, wrt, req); err != nil {
					wrt.WriteHeader(http.StatusInternalServerError)
					wrt.Write([]byte(err.Error()))
				}
			}
		}
	}
}

// Get get
func (p *Mux) Get(pattern string, handlers ...interface{}) {
	p.add(http.MethodGet, pattern, handlers...)
}

// Post post
func (p *Mux) Post(pattern string, handlers ...interface{}) {
	p.add(http.MethodGet, pattern, handlers...)
}

// Put put
func (p *Mux) Put(pattern string, handlers ...interface{}) {
	p.add(http.MethodGet, pattern, handlers...)
}

// Patch patch
func (p *Mux) Patch(pattern string, handlers ...interface{}) {
	p.add(http.MethodGet, pattern, handlers...)
}

// Delete delte
func (p *Mux) Delete(pattern string, handlers ...interface{}) {
	p.add(http.MethodGet, pattern, handlers...)
}

func (p *Mux) walk(f func(method, pattern string, handlers ...interface{}) error) error {
	for _, r := range p.routes {
		if err := f(r.method, r.pattern, r.handlers...); err != nil {
			return err
		}
	}
	return nil
}

func (p *Mux) add(method, pattern string, handlers ...interface{}) {
	p.routes = append(
		p.routes, route{
			method:   method,
			pattern:  pattern,
			handlers: handlers,
		},
	)
}
