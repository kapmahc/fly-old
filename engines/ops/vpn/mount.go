package vpn

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/vpn/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/users", auth.HTML(p.indexUsers))
	ag.GET("/users/new", auth.HTML(p.createUser))
	ag.POST("/users/new", auth.HTML(p.createUser))
	ag.GET("/users/edit/:id", auth.HTML(p.updateUser))
	ag.POST("/users/edit/:id", auth.HTML(p.updateUser))
	ag.GET("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.GET("/logs", auth.HTML(p.indexLogs))
	ag.GET("/readme", auth.HTML(p.getReadme))

	rt.GET("/users/change-password/:id", auth.HTML(p.changeUserPassword))
	rt.POST("/users/change-password/:id", auth.HTML(p.changeUserPassword))

	api := rt.Group("/vpn/api", p.TokenMiddleware)
	api.POST("/auth", web.JSON(p.apiAuth))
	api.POST("/connect", web.JSON(p.apiConnect))
	api.POST("/disconnect", web.JSON(p.apiDisconnect))
}
