package mail

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/ops/mail", p.Jwt.MustAdminMiddleware)

	ag.GET("/users", web.JSON(p.indexUsers))
	ag.GET("/users/new", web.JSON(p.createUser))
	ag.POST("/users/new", web.JSON(p.createUser))
	ag.GET("/users/edit/:id", web.JSON(p.updateUser))
	ag.POST("/users/edit/:id", web.JSON(p.updateUser))
	ag.GET("/users/reset-password/:id", web.JSON(p.resetUserPassword))
	ag.POST("/users/reset-password/:id", web.JSON(p.resetUserPassword))
	ag.DELETE("/users/:id", web.JSON(p.destroyUser))

	rt.GET("/ops/mail/users/change-password", web.JSON(p.changeUserPassword))
	rt.POST("/ops/mail/users/change-password", web.JSON(p.changeUserPassword))

	ag.GET("/domains", web.JSON(p.indexDomains))
	ag.GET("/domains/new", web.JSON(p.createDomain))
	ag.POST("/domains/new", web.JSON(p.createDomain))
	ag.GET("/domains/edit/:id", web.JSON(p.updateDomain))
	ag.POST("/domains/edit/:id", web.JSON(p.updateDomain))
	ag.DELETE("/domains/:id", web.JSON(p.destroyDomain))

	ag.GET("/aliases", web.JSON(p.indexAliases))
	ag.GET("/aliases/new", web.JSON(p.createAlias))
	ag.POST("/aliases/new", web.JSON(p.createAlias))
	ag.GET("/aliases/edit/:id", web.JSON(p.updateAlias))
	ag.POST("/aliases/edit/:id", web.JSON(p.updateAlias))
	ag.DELETE("/aliases/:id", web.JSON(p.destroyAlias))

	ag.GET("/readme", web.JSON(p.getReadme))
}
