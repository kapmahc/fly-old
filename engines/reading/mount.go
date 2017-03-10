package reading

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/reading")
	rg.GET("/books", p.indexBooks)
	rg.GET("/books/:id", p.showBook)
	rg.GET("/pages/:id/*href", p.showPage)
}
