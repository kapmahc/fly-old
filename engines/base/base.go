package base

import "github.com/astaxie/beego"

// Controller base
type Controller struct {
	beego.Controller
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.Layout = "application.html"
}
