package reading

import (
	"net/http"
)

// Mount web mount points
func (p *Engine) Mount() {

	rg := p.Router.PathPrefix("/reading").Subrouter()
	rg.HandleFunc("/", p.indexBooks).Methods(http.MethodGet).Name("reading.engine.home")
	rg.HandleFunc("/books", p.indexBooks).Methods(http.MethodGet).Name("reading.books.index")
	rg.HandleFunc("/books/{id:[0-9]+}", p.showBook).Methods(http.MethodGet).Name("reading.book.show")
	rg.HandleFunc(`/books/{id:[0-9]+}/{name:[\w\/\.-]+}`, p.showBookPage).Methods(http.MethodGet).Name("reading.book.page")

}
