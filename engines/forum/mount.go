package forum

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	fg := rt.Group("/forum")
	fg.GET("/articles", p.indexArticles)
	fg.POST("/articles", p.Jwt.MustSignInMiddleware, p.createArticle)
	fg.GET("/articles/:id", p.showArticle)
	fg.POST("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, p.updateArticle)
	fg.DELETE("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, p.destroyArticle)

	fg.GET("/tags", p.indexTags)
	fg.POST("/tags", p.Jwt.MustAdminMiddleware, p.createTag)
	fg.GET("/tags/:id", p.showTag)
	fg.POST("/tags/:id", p.Jwt.MustAdminMiddleware, p.updateTag)
	fg.DELETE("/tags/:id", p.Jwt.MustAdminMiddleware, p.destroyTag)

	fg.GET("/comments", p.indexComments)
	fg.POST("/comments", p.Jwt.MustSignInMiddleware, p.createComment)
	fg.POST("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, p.updateComment)
	fg.DELETE("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, p.destroyComment)

}
