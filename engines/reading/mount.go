package reading

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mount web mount points
func (p *Engine) Mount(rt *mux.Router) {

	rg := rt.PathPrefix("/reading").Subrouter()
	rg.HandleFunc("/books", p.indexBooks).Methods(http.MethodGet).Name("reading.books.index")
	rg.HandleFunc("/books/{id:[0-9]+}", p.showBook).Methods(http.MethodGet).Name("reading.book.show")
	rg.HandleFunc(`/books/{id:[0-9]+}/{name:[\w\/\.]+}`, p.showBookPage).Methods(http.MethodGet).Name("reading.book.page")

}
