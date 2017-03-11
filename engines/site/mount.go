package site

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	rt.GET("/locales/:lang", p.getLocales)
	rt.GET("/site/info", p.getSiteInfo)

	rt.GET("/notices", p.indexNotices)
	rt.POST("/notices", p.Jwt.MustAdminMiddleware, p.createNotice)
	rt.POST("/notices/:id", p.Jwt.MustAdminMiddleware, p.updateNotice)
	rt.DELETE("/notices/:id", p.Jwt.MustAdminMiddleware, p.destroyNotice)

	rt.GET("/leave-words", p.Jwt.MustAdminMiddleware, p.indexLeaveWords)
	rt.POST("/leave-words", p.createLeaveWord)
	rt.DELETE("/leave-words/:id", p.Jwt.MustAdminMiddleware, p.destroyLeaveWord)

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)
	ag.GET("/locales", p.getAdminLocales)
	ag.POST("/locales", p.postAdminLocales)
	ag.DELETE("/locales/:id", p.deleteAdminLocales)
	ag.GET("/users", p.getAdminUsers)

	asg := ag.Group("/site")
	asg.GET("/status", p.getAdminSiteStatus)
	asg.POST("/info", p.postAdminSiteInfo)
	asg.POST("/author", p.postAdminSiteAuthor)
	asg.GET("/seo", p.getAdminSiteSeo)
	asg.POST("/seo", p.postAdminSiteSeo)
	asg.GET("/smtp", p.getAdminSiteSMTP)
	asg.POST("/smtp", p.postAdminSiteSMTP)

}
