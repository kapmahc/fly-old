package blog

// Mount web mount points
func (p *Engine) Mount() {
	bg := p.Mux.Group("/blog")
	bg.Get("blog.index", "/", p.indexPosts)
	bg.Get("blog.show", `/{name:[\w\/\.-]+}`, p.showPost)

}
