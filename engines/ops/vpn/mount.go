package vpn

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/vpn", p.Jwt.MustAdminMiddleware)
	ag.GET("/users", auth.HTML(p.indexUsers))
	ag.GET("/users/new", auth.HTML(p.createUser))
	ag.POST("/users/new", auth.HTML(p.createUser))
	ag.GET("/users/edit/:id", auth.HTML(p.updateUser))
	ag.POST("/users/edit/:id", auth.HTML(p.updateUser))
	ag.GET("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", auth.HTML(p.resetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	ag.GET("/logs", auth.HTML(p.indexLogs))

	ag.GET("/readme", auth.HTML(p.getReadme))

	rt.GET("/vpn/users/change-password", auth.HTML(p.changeUserPassword))
	rt.POST("/vpn/users/change-password", auth.HTML(p.changeUserPassword))

	api := rt.Group("/vpn/api", p.tokenMiddleware)
	api.POST("/auth", web.JSON(p.apiAuth))
	api.POST("/connect", web.JSON(p.apiConnect))
	api.POST("/disconnect", web.JSON(p.apiDisconnect))
}
