package blog

import (
	"net/http"
)

// Mount web mount points
func (p *Engine) Mount() {
	bg := p.Router.PathPrefix("/blog").Subrouter()
	bg.HandleFunc("/", p.indexPosts).Methods(http.MethodGet).Name("blog.engine.home")
	bg.HandleFunc(`/{name:[\w\/\.-]+}`, p.showPost).Methods(http.MethodGet).Name("blog.show")

}
