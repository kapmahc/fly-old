package site

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/install", auth.HTML(p.formInstall))
	rt.POST("/install", auth.HTML(p.formInstall))
	rt.GET("/dashboard", p.Jwt.MustSignInMiddleware, auth.HTML(p.getDashboard))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/locales", auth.HTML(p.getAdminLocales))
	ag.GET("/locales/edit", auth.HTML(p.formAdminLocales))
	ag.POST("/locales/edit", auth.HTML(p.formAdminLocales))
	ag.DELETE("/locales/:id", web.JSON(p.deleteAdminLocales))
	ag.GET("/users", auth.HTML(p.getAdminUsers))

	asg := ag.Group("/site")
	asg.GET("/status", auth.HTML(p.getAdminSiteStatus))
	asg.GET("/info", auth.HTML(p.formAdminSiteInfo))
	asg.POST("/info", auth.HTML(p.formAdminSiteInfo))
	asg.GET("/author", auth.HTML(p.formAdminSiteAuthor))
	asg.POST("/author", auth.HTML(p.formAdminSiteAuthor))
	asg.GET("/seo", auth.HTML(p.formAdminSiteSeo))
	asg.POST("/seo", auth.HTML(p.formAdminSiteSeo))
	asg.GET("/smtp", auth.HTML(p.formAdminSiteSMTP))
	asg.POST("/smtp", auth.HTML(p.formAdminSiteSMTP))
	// ----------------

	rt.GET("/notices", web.JSON(p.indexNotices))
	rt.POST("/notices", p.Jwt.MustAdminMiddleware, web.JSON(p.createNotice))
	rt.POST("/notices/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.updateNotice))
	rt.DELETE("/notices/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyNotice))

	rt.GET("/leave-words", p.Jwt.MustAdminMiddleware, web.JSON(p.indexLeaveWords))
	rt.POST("/leave-words", web.JSON(p.createLeaveWord))
	rt.DELETE("/leave-words/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyLeaveWord))

}
