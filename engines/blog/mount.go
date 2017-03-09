package blog

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/blogs", p.indexPosts)
	rt.GET("/blog/*href", p.showPost)
}
