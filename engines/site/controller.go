package site

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/kapmahc/fly/engines/auth"
)

// Controller controller
type Controller struct {
	auth.UserController
}

// GetHome index
// @router / [get]
func (p *Controller) GetHome() {
	u := fmt.Sprintf("%s.Controller.GetHome", beego.AppConfig.String("homeengine"))
	if p.dbEmpty() {
		u = "site.Controller.GetInstall"
	}
	p.Redirect(p.URLFor(u), http.StatusFound)
}

// GetDashboard dashboard
// @router /dashboard [get]
func (p *Controller) GetDashboard() {
	p.HTML(p.T("site.dashboard.title"), "site/dashboard.html")
}
