package site

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/engines/base"
)

// Controller controller
type Controller struct {
	base.Controller
}

// GetHome index
// @router / [get]
func (p *Controller) GetHome() {
	o := orm.NewOrm()
	count, err := o.QueryTable(&auth.User{}).Count()
	if err != nil {
		beego.Error(err)
		p.Abort("500")
	}
	if count == 0 {
		p.Redirect(p.URLFor("site.Controller.GetInstall"), http.StatusFound)
		return
	}
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
