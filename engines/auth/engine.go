package auth

import "github.com/astaxie/beego"

// Controller controller
type Controller struct {
	beego.Controller
}

// Index index
// @router / [get]
func (c *Controller) Index() {

}

// SignIn sign in
// @router /sign-in [get]
func (c *Controller) SignIn() {

}
