package reading

// Mount web mount points
func (p *Engine) Mount() {

	rg := p.Mux.Group("/reading")
	rg.Get("reading.engine.home", "/", p.indexBooks)
	rg.Get("reading.books.index", "/books", p.indexBooks)
	rg.Get("reading.book.show", "/books/{id:[0-9]+}", p.showBook)
	rg.Get("reading.book.page", `/books/{id:[0-9]+}/{name:[\w\/\.-]+}`, p.showBookPage)

}
