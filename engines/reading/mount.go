package reading

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/reading")
	rg.GET("/books", web.JSON(p.indexBooks))
	rg.GET("/books/:id", web.JSON(p.showBook))
	rg.GET("/pages/:id/*href", p.showPage)
}
