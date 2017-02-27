package site

// Mount web mount points
func (p *Engine) Mount() {
	p.Mux.Get("site.engine.home", "/", p.home)
	p.Mux.Get("site.notices.index", "/notices", p.indexNotices)
}
