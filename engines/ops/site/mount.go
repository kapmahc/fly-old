package site

import "net/http"

// Mount web mount points
func (p *Engine) Mount() {
	p.Router.HandleFunc("/", p.home).Methods(http.MethodGet).Name("site.engine.home")
	p.Router.HandleFunc("/install", p.install).Methods(http.MethodGet, http.MethodPost).Name("site.install")

}
