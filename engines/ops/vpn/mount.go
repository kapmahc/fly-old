package vpn

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.POST("/ops/vpn/users/change-password", web.JSON(p.postChangeUserPassword))

	ag := rt.Group("/ops/vpn", p.Jwt.MustAdminMiddleware)
	ag.GET("/users", web.JSON(p.indexUsers))
	ag.POST("/users", web.JSON(p.createUser))
	ag.POST("/users/info/:id", web.JSON(p.updateUser))
	ag.POST("/users/reset-password/:id", web.JSON(p.postResetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	ag.GET("/logs", web.JSON(p.indexLogs))

	ag.GET("/readme", p.getReadme)

	api := rt.Group("/ops/vpn", p.tokenMiddleware)
	api.POST("/auth", web.JSON(p.apiAuth))
	api.POST("/connect", web.JSON(p.apiConnect))
	api.POST("/disconnect", web.JSON(p.apiDisconnect))
}
