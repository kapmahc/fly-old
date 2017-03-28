package forum

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {

	fg := rt.Group("/forum")
	fg.GET("/tags", web.JSON(p.indexTags))
	fg.POST("/tags", web.JSON(p.createTag))
	fg.GET("/tags/:id", web.JSON(p.showTag))
	fg.POST("/tags/:id", web.JSON(p.updateTag))
	fg.DELETE("/tags/:id", web.JSON(p.destroyTag))

	fg.GET("/articles", web.JSON(p.indexArticles))
	fg.POST("/articles", p.Jwt.MustSignInMiddleware, web.JSON(p.createArticle))
	fg.GET("/articles/:id", web.JSON(p.showArticle))
	fg.POST("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.updateArticle))
	fg.DELETE("/articles/:id", p.Jwt.MustSignInMiddleware, p.canEditArticle, web.JSON(p.destroyArticle))

	fg.GET("/comments", web.JSON(p.indexComments))
	fg.POST("/comments", p.Jwt.MustSignInMiddleware, web.JSON(p.createComment))
	fg.POST("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.updateComment))
	fg.DELETE("/comments/:id", p.Jwt.MustSignInMiddleware, p.canEditComment, web.JSON(p.destroyComment))

	mg := fg.Group("/my", p.Jwt.MustSignInMiddleware)
	mg.GET("/articles", web.JSON(p.myArticles))
	mg.GET("/comments", web.JSON(p.myComments))
}
