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

	}, p.Jwt.MustAdminMiddleware, p.Layout.Dashboard)
}
