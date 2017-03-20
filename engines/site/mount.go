package site

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/", auth.HTML(p.getHome))
	rt.GET("/install", auth.HTML(p.formInstall))
	rt.POST("/install", auth.HTML(p.formInstall))
	rt.GET("/dashboard", p.Jwt.MustSignInMiddleware, auth.HTML(p.getDashboard))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)

	ag.GET("/locales", auth.HTML(p.getAdminLocales))
	ag.GET("/locales/edit", auth.HTML(p.formAdminLocales))
	ag.POST("/locales/edit", auth.HTML(p.formAdminLocales))
	ag.DELETE("/locales/:id", web.JSON(p.deleteAdminLocales))

	ag.GET("/users", auth.HTML(p.getAdminUsers))

	ag.GET("/links", auth.HTML(p.indexAdminLinks))
	ag.GET("/links/new", auth.HTML(p.createAdminLink))
	ag.POST("/links/new", auth.HTML(p.createAdminLink))
	ag.GET("/links/edit/:id", auth.HTML(p.updateAdminLink))
	ag.POST("/links/edit/:id", auth.HTML(p.updateAdminLink))
	ag.DELETE("/links/:id", web.JSON(p.destroyAdminLink))

	ag.GET("/pages", auth.HTML(p.indexAdminPages))
	ag.GET("/pages/new", auth.HTML(p.createAdminPage))
	ag.POST("/pages/new", auth.HTML(p.createAdminPage))
	ag.GET("/pages/edit/:id", auth.HTML(p.updateAdminPage))
	ag.POST("/pages/edit/:id", auth.HTML(p.updateAdminPage))
	ag.DELETE("/pages/:id", web.JSON(p.destroyAdminPage))

	ag.GET("/notices", auth.HTML(p.indexAdminNotices))
	ag.GET("/notices/new", auth.HTML(p.createNotice))
	ag.POST("/notices/new", auth.HTML(p.createNotice))
	ag.GET("/notices/edit/:id", auth.HTML(p.updateNotice))
	ag.POST("/notices/edit/:id", auth.HTML(p.updateNotice))
	ag.DELETE("/notices/:id", web.JSON(p.destroyNotice))

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

	rt.GET("/notices", auth.HTML(p.indexNotices))

	rt.GET("/leave-words", p.Jwt.MustAdminMiddleware, auth.HTML(p.indexLeaveWords))
	rt.GET("/leave-words/new", auth.HTML(p.createLeaveWord))
	rt.POST("/leave-words/new", auth.HTML(p.createLeaveWord))
	rt.DELETE("/leave-words/:id", p.Jwt.MustAdminMiddleware, web.JSON(p.destroyLeaveWord))

}
