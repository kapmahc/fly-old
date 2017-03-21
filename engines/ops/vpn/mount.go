package vpn

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/vpn/api", p.Jwt.Middleware)
	ag.POST("/auth", web.JSON(p.apiAuth))
	ag.POST("/connect", web.JSON(p.apiConnect))
	ag.POST("/disconnect", web.JSON(p.apiDisconnect))
}
