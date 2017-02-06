package site

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
)

// Controller controller
type Controller struct {
	auth.UserController
	// base.Controller
}

// GetHome index
// @router / [get]
func (p *Controller) GetHome() {
	if p.dbEmpty() {
		p.Redirect(p.URLFor("site.Controller.GetInstall"), http.StatusFound)
		return
	}

	p.HTML(p.T("site.home.title"), "site/home.html")
}

// GetDashboard dashboard
// @router /dashboard [get]
func (p *Controller) GetDashboard() {
	p.HTML(p.T("site.dashboard.title"), "site/dashboard.html")
}
