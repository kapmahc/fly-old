package forum

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/forum/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/tags", web.JSON(p.indexAdminTags))
	ag.GET("/tags/new", web.JSON(p.createTag))
	ag.POST("/tags/new", web.JSON(p.createTag))
	ag.GET("/tags/edit/:id", web.JSON(p.updateTag))
	ag.POST("/tags/edit/:id", web.JSON(p.updateTag))
	ag.DELETE("/tags/:id", web.JSON(p.destroyTag))

	fg := rt.Group("/forum")
	fg.GET("/articles", web.JSON(p.indexArticles))
	fg.GET("/articles/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createArticle))
	fg.POST("/articles/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createArticle))
	fg.GET("/articles/show/:id", web.JSON(p.showArticle))
	fg.GET("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.updateArticle))
	fg.POST("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.updateArticle))
	fg.DELETE("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.destroyArticle))

	fg.GET("/tags", web.JSON(p.indexTags))
	fg.GET("/tags/show/:id", web.JSON(p.showTag))

	fg.GET("/comments", web.JSON(p.indexComments))
	fg.GET("/comments/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createComment))
	fg.POST("/comments/new", p.Jwt.MustSignInMiddleware, web.JSON(p.createComment))
	fg.GET("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.updateComment))
	fg.POST("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.updateComment))
	fg.DELETE("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.destroyComment))

	fg.GET("/articles/my", p.Jwt.MustSignInMiddleware, web.JSON(p.myArticles))
	fg.GET("/comments/my", p.Jwt.MustSignInMiddleware, web.JSON(p.myComments))
}
