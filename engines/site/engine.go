package site

import "github.com/astaxie/beego"

// Controller controller
type Controller struct {
	beego.Controller
}

// GetHome get home
// @router / [get]
func (c *Controller) GetHome() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
