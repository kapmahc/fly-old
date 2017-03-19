package blog

import (
	"github.com/kapmahc/fly/engines/auth"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/blogs", auth.HTML(p.indexPosts))
	rt.GET("/blog/*href", auth.HTML(p.showPost))
}
