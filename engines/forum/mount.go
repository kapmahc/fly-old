package forum

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	fg := rt.Group("/forum")
	fg.GET("/articles", auth.HTML(p.indexArticles))
	fg.GET("/articles/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createArticle))
	fg.POST("/articles/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createArticle))
	fg.GET("/articles/show/:id", auth.HTML(p.showArticle))
	fg.GET("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, auth.HTML(p.updateArticle))
	fg.POST("/articles/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, auth.HTML(p.updateArticle))
	fg.DELETE("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.destroyArticle))

	fg.GET("/tags", auth.HTML(p.indexTags))
	fg.GET("/tags/new", p.Jwt.MustAdminMiddleware, auth.HTML(p.createTag))
	fg.POST("/tags/new", p.Jwt.MustAdminMiddleware, auth.HTML(p.createTag))
	fg.GET("/tags/show/:id", auth.HTML(p.showTag))
	fg.GET("/tags/edit/:id", p.Jwt.MustAdminMiddleware, auth.HTML(p.updateTag))
	fg.POST("/tags/edit/:id", p.Jwt.MustAdminMiddleware, auth.HTML(p.updateTag))
	fg.DELETE("/tags/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyTag))

	fg.GET("/comments", auth.HTML(p.indexComments))
	fg.GET("/comments/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createComment))
	fg.POST("/comments/new", p.Jwt.MustSignInMiddleware, auth.HTML(p.createComment))
	fg.GET("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, auth.HTML(p.updateComment))
	fg.POST("/comments/edit/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, auth.HTML(p.updateComment))
	fg.DELETE("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.destroyComment))

	fg.GET("/articles/my", p.Jwt.MustSignInMiddleware, auth.HTML(p.myArticles))
	fg.GET("/comments/my", p.Jwt.MustSignInMiddleware, auth.HTML(p.myComments))
}
