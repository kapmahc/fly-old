package vpn

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/ops/vpn", p.Jwt.MustAdminMiddleware)
	ag.GET("/users", web.JSON(p.indexUsers))
	ag.GET("/users/new", web.JSON(p.createUser))
	ag.POST("/users/new", web.JSON(p.createUser))
	ag.GET("/users/edit/:id", web.JSON(p.updateUser))
	ag.POST("/users/edit/:id", web.JSON(p.updateUser))
	ag.GET("/users/reset-password/:id", web.JSON(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", web.JSON(p.resetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	ag.GET("/logs", web.JSON(p.indexLogs))

	ag.GET("/readme", web.JSON(p.getReadme))

	rt.GET("/ops/vpn/users/change-password", web.JSON(p.changeUserPassword))
	rt.POST("/ops/vpn/users/change-password", web.JSON(p.changeUserPassword))

	api := rt.Group("/ops/vpn/api", p.tokenMiddleware)
	api.POST("/auth", web.JSON(p.apiAuth))
	api.POST("/connect", web.JSON(p.apiConnect))
	api.POST("/disconnect", web.JSON(p.apiDisconnect))
}
