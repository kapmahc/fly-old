package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Mux mux
type Mux struct {
	Router *mux.Router `inject:""`
}

// Status print routes
func (p *Mux) Status(w io.Writer) {
	tpl := "%-32s %s\n"
	fmt.Fprintf(w, tpl, "NAME", "PATH")
	if err := p.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pat, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		hnd := route.GetHandler()
		if hnd == nil {
			return nil
		}
		fmt.Fprintf(
			w,
			tpl,
			route.GetName(),
			pat,
			// runtime.FuncForPC(reflect.ValueOf(route.GetHandler()).Pointer()).Name(),
		)
		return nil
	}); err != nil {

	}
}

// Crud crud
func (p *Mux) Crud(name, pat string, index, new, show, edit, destroy http.HandlerFunc) {
	if index != nil {
		p.Get(name+".index", pat+"/", index)
	}
	if new != nil {
		p.Form(name+".new", pat+"/new", new)
	}
	if edit != nil {
		p.Form(name+".edit", pat+"/edit/{id:[0-9]+}", edit)
	}
	if show != nil {
		p.Get(name+".show", pat+"/{id:[0-9]+}", show)
	}
	if destroy != nil {
		p.Delete(name+".destroy", pat+"/{id:[0-9]+}", destroy)
	}

}

// Group group
func (p *Mux) Group(pat string) *Mux {
	return &Mux{Router: p.Router.PathPrefix(pat).Subrouter()}
}

// Get get
func (p *Mux) Get(name string, pat string, hnd http.HandlerFunc) {
	p.Router.HandleFunc(pat, hnd).Methods(http.MethodGet).Name(name)
}

// Post post
func (p *Mux) Post(name string, pat string, hnd http.HandlerFunc) {
	p.Router.HandleFunc(pat, hnd).Methods(http.MethodPost).Name(name)
}

// Form post form
func (p *Mux) Form(name string, pat string, hnd http.HandlerFunc) {
	p.Router.HandleFunc(pat, hnd).Methods(http.MethodGet, http.MethodPost).Name(name)
}

// Delete delete
func (p *Mux) Delete(name string, pat string, hnd http.HandlerFunc) {
	p.Router.HandleFunc(pat, hnd).Methods(http.MethodDelete).Name(name)
}
