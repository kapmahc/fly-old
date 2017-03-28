package mail

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.POST("/ops/mail/users/change-password", web.JSON(p.postChangeUserPassword))

	ag := rt.Group("/ops/mail", p.Jwt.MustAdminMiddleware)

	ag.GET("/users", web.JSON(p.indexUsers))
	ag.POST("/users", web.JSON(p.createUser))
	ag.POST("/users/info/:id", web.JSON(p.updateUser))
	ag.POST("/users/reset-password/:id", web.JSON(p.postResetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	ag.GET("/domains", web.JSON(p.indexDomains))
	ag.POST("/domains", web.JSON(p.createDomain))
	ag.POST("/domains/:id", web.JSON(p.updateDomain))
	ag.DELETE("/domains/:id", web.JSON(p.destroyDomain))

	ag.GET("/aliases", web.JSON(p.indexAliases))
	ag.POST("/aliases", web.JSON(p.createAlias))
	ag.POST("/aliases/:id", web.JSON(p.updateAlias))
	ag.DELETE("/aliases/:id", web.JSON(p.destroyAlias))

	ag.GET("/readme", p.getReadme)
}
