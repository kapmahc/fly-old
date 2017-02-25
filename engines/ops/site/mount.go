package site

import "net/http"

// Mount web mount points
func (p *Engine) Mount() {
	rt := p.Router
	rt.HandleFunc("/", p.home).Methods(http.MethodGet).Name("site.engine.home")
	rt.HandleFunc("/notices", p.indexNotices).Methods(http.MethodGet, http.MethodPost).Name("site.notices.index")

}
