package site

import "github.com/kapmahc/sky"

// Mount web mount points
func (p *Engine) Mount(rt *sky.Router) {
	rt.Get("site.install", "/install", p.mustDatabaseEmpty, p.Layout.Application, p.getInstall)
	rt.Post("site.install", "/install", p.mustDatabaseEmpty, p.postInstall)

	rt.Get("site.dashboard", "/dashboard", p.Jwt.MustSignInMiddleware, p.Layout.Dashboard, p.getDashboard)
}
