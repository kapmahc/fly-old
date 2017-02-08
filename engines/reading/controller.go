package reading

import "github.com/kapmahc/fly/engines/auth"

// Controller controller
type Controller struct {
	auth.UserController
}

// GetHome get home
// @router / [get]
func (p *Controller) GetHome() {
	p.IndexBooks()
}
