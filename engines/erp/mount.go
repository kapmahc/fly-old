package erp

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rg := rt.Group("/erp", p.Jwt.MustSignInMiddleware, p.mustSellerMiddleware)
	rg.GET("/tags", auth.HTML(p.indexTags))
	rg.GET("/tags/new", auth.HTML(p.createTag))
	rg.POST("/tags/new", auth.HTML(p.createTag))
	rg.GET("/tags/edit/:id", auth.HTML(p.updateTag))
	rg.POST("/tags/edit/:id", auth.HTML(p.updateTag))
	rg.DELETE("/tags/:id", web.JSON(p.destroyTag))
}
