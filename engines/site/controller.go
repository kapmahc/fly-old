package site

import "github.com/kapmahc/fly/engines/base"

// Controller controller
type Controller struct {
	base.Controller
}

// GetHome index
// @router / [get]
func (p *Controller) GetHome() {
	p.Data["title"] = p.T("site.home.title")
	p.TplName = "site/home.html"
}

// GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	p.Data["title"] = p.T("site.install.title")
	p.TplName = "site/install.html"
}

// PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
	p.TplName = "site/install.html"
}

// Dashboard get index
// @router /dashboard [get]
func (p *Controller) Dashboard() {
	p.TplName = "site/dashboard.html"
}
