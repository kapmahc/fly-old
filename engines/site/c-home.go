package site

import (
	"net/http"

	"github.com/kapmahc/sky"
)

func (p *Engine) getHome(wrt http.ResponseWriter, req *http.Request) {
	var home string
	if err := p.Settings.Get("site.home", &home); err != nil {
		home = "site.notices.index"
	}
	rt := p.Router.Get(home)
	if rt != nil {
		rt.GetHandler().ServeHTTP(wrt, req)
	}
}

func (p *Engine) getDashboard(c *sky.Context) error {
	c.HTML(http.StatusOK, "site/dashboard", c.Get(sky.DATA))
	return nil
}
