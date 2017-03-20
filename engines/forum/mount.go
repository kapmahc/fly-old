package forum

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/forum/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/tags", auth.HTML(p.indexAdminTags))
	ag.GET("/tags/new", auth.HTML(p.createTag))
	ag.POST("/tags/new", auth.HTML(p.createTag))
	ag.GET("/tags/edit/:id", auth.HTML(p.updateTag))
	ag.POST("/tags/edit/:id", auth.HTML(p.updateTag))
	ag.DELETE("/tags/:id", web.JSON(p.destroyTag))

	fg := rt.Group("/forum")
	fg.GET("/articles", auth.HTML(p.indexArticles))
	fg.GET("/articles/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createArticle))
	fg.POST("/articles/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createArticle))
	fg.GET("/articles/show/:id", auth.HTML(p.showArticle))
	fg.GET("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, auth.HTML(p.updateArticle))
	fg.POST("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, auth.HTML(p.updateArticle))
	fg.DELETE("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.destroyArticle))

	fg.GET("/tags", auth.HTML(p.indexTags))
	fg.GET("/tags/show/:id", auth.HTML(p.showTag))

	fg.GET("/comments", auth.HTML(p.indexComments))
	fg.GET("/comments/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createComment))
	fg.POST("/comments/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createComment))
	fg.GET("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, auth.HTML(p.updateComment))
	fg.POST("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, auth.HTML(p.updateComment))
	fg.DELETE("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.destroyComment))

	fg.GET("/articles/my", p.Jwt.MustSignInMiddleware, auth.HTML(p.myArticles))
	fg.GET("/comments/my", p.Jwt.MustSignInMiddleware, auth.HTML(p.myComments))
}
