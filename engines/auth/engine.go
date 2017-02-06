package auth

import "github.com/kapmahc/fly/engines/base"

// Controller controller
type Controller struct {
	base.Controller
}

// Index index
// @router / [get]
func (c *Controller) Index() {

}

// SignIn sign in
// @router /sign-in [get]
func (c *Controller) SignIn() {

}
