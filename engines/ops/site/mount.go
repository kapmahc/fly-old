package site

// Mount web mount points
func (p *Engine) Mount() {
	p.Mux.Get("site.home", "/", p.home)
	p.Mux.Get("site.dashboard", "/dashboard", p.dashboard)

	p.Mux.Get("site.notices.index", "/notices", p.indexNotices)

}
