package reading

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/reading")
	rg.GET("/books", auth.HTML(p.indexBooks))
	rg.GET("/books/:id", auth.HTML(p.showBook))
	rg.GET("/pages/:id/*href", p.showPage)
	rg.GET("/dict", auth.HTML(p.formDict))
	rg.POST("/dict", auth.HTML(p.formDict))

	ag := rg.Group("/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/status", auth.HTML(p.getAdminStatus))
	ag.GET("/books", auth.HTML(p.indexAdminBooks))
	ag.DELETE("/books/:id", web.JSON(p.destroyAdminBook))
}
