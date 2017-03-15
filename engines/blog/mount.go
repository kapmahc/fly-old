package blog

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/blogs", web.JSON(p.indexPosts))
	rt.GET("/blog/*href", web.JSON(p.showPost))
}
