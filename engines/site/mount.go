package site

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/install", auth.HTML(p.getInstall))
	rt.POST("/install", auth.HTML(p.postInstall))
	// ----------------
	rt.GET("/locales/:lang", web.JSON(p.getLocales))
	rt.GET("/site/info", web.JSON(p.getSiteInfo))

	rt.GET("/notices", web.JSON(p.indexNotices))
	rt.POST("/notices", p.Jwt.MustAdminMiddleware, web.JSON(p.createNotice))
	rt.POST("/notices/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.updateNotice))
	rt.DELETE("/notices/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyNotice))

	rt.GET("/leave-words", p.Jwt.MustAdminMiddleware, web.JSON(p.indexLeaveWords))
	rt.POST("/leave-words", web.JSON(p.createLeaveWord))
	rt.DELETE("/leave-words/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyLeaveWord))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/locales", web.JSON(p.getAdminLocales))
	ag.POST("/locales", web.JSON(p.postAdminLocales))
	ag.DELETE("/locales/:id", web.JSON(p.deleteAdminLocales))
	ag.GET("/users", web.JSON(p.getAdminUsers))

	asg := ag.Group("/site")
	asg.GET("/status", web.JSON(p.getAdminSiteStatus))
	asg.POST("/info", web.JSON(p.postAdminSiteInfo))
	asg.POST("/author", web.JSON(p.postAdminSiteAuthor))
	asg.GET("/seo", web.JSON(p.getAdminSiteSeo))
	asg.POST("/seo", web.JSON(p.postAdminSiteSeo))
	asg.GET("/smtp", web.JSON(p.getAdminSiteSMTP))
	asg.POST("/smtp", web.JSON(p.postAdminSiteSMTP))

}
