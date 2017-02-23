package blog

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mount web mount points
func (p *Engine) Mount(rt *mux.Router) {
	bg := rt.PathPrefix("/blog").Subrouter()
	bg.HandleFunc("/", p.indexPosts).Methods(http.MethodGet).Name("blog.home")
	bg.HandleFunc(`/{name:[\w\/\.-]+}`, p.showPost).Methods(http.MethodGet).Name("blog.show")

}
