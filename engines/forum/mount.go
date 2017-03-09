package forum

// Mount web mount points
func (p *Engine) Mount() {
	p.Mux.Crud("forum.articles", "/forum/articles", p.indexArticles, p.newArticle, p.showArticle, p.editArticle, p.destroyArticle)
	p.Mux.Crud("forum.tags", "/forum/tags", p.indexTags, p.newTag, p.showTag, p.editTag, p.destroyTag)
	p.Mux.Crud("forum.comments", "/forum/comments", p.indexComments, p.newComment, nil, p.editComment, p.destroyComment)

	dg := p.Mux.Group("/forum/dashboard")
	dg.Get("forum.dashboard.articles", "/articles", p.dashboardArticles)
	dg.Get("forum.dashboard.tags", "/tags", p.dashboardTags)
	dg.Get("forum.dashboard.comments", "/comments", p.dashboardComments)
}
