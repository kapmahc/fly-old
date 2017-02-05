package shop

import "github.com/astaxie/beego"

// Controller controller
type Controller struct {
	beego.Controller
}

// GetHome get home
// @router / [get]
func (c *Controller) GetHome() {
	c.TplName = "shop/index.tpl"
}
