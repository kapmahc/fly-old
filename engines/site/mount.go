package site

import (
	"net/http"

	"github.com/kapmahc/sky"
)

// Mount web mount points
func (p *Engine) Mount(rt *sky.Router) {
	p.Router.HandleFunc("/", p.getHome).Methods(http.MethodGet)

	rt.Get("site.install", "/install", p.mustDatabaseEmpty, p.Layout.Application, p.getInstall)
	rt.Post("site.install", "/install", p.mustDatabaseEmpty, p.postInstall)

	rt.Get("site.dashboard", "/dashboard", p.Jwt.MustSignInMiddleware, p.Layout.Dashboard, p.getDashboard)

	rt.Group("/admin", func(r *sky.Router) {
		r.Get("site.admin.status", "/site/status", p.getAdminSiteStatus)

		r.Get("site.admin.info", "/site/info", p.getAdminSiteInfo)
		r.Post("site.admin.info", "/site/info", p.postAdminSiteInfo)
		r.Get("site.admin.author", "/site/author", p.getAdminSiteAuthor)
		r.Post("site.admin.author", "/site/author", p.postAdminSiteAuthor)
		r.Get("site.admin.seo", "/site/seo", p.getAdminSiteSeo)
		r.Post("site.admin.seo", "/site/seo", p.postAdminSiteSeo)
		r.Get("site.admin.smtp", "/site/smtp", p.getAdminSiteSMTP)
		r.Post("site.admin.smtp", "/site/smtp", p.postAdminSiteSMTP)

	}, p.Jwt.MustAdminMiddleware, p.Layout.Dashboard)
}
