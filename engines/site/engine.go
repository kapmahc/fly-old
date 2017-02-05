package site

import "github.com/kapmahc/fly/engines/base"

// Controller controller
type Controller struct {
	base.Controller
}

// Home index
// @router / [get]
func (p *Controller) Home() {
	p.TplName = "site/home.html"
}

// Dashboard get index
// @router /dashboard [get]
func (p *Controller) Dashboard() {
	p.TplName = "site/dashboard.html"
}
