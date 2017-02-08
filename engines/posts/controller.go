package posts

import "github.com/kapmahc/fly/engines/auth"

// Controller controller
type Controller struct {
	auth.UserController
}

// GetHome home
// @router / [get]
func (p *Controller) GetHome() {
	p.HTML(p.T("posts.home.title"), "posts/home.html")
}
