package site

import (
	"net/http"

	"github.com/kapmahc/sky"
)

func (p *Engine) getDashboard(c *sky.Context) error {
	c.HTML(http.StatusOK, "site/dashboard", c.Get(sky.DATA))
	return nil
}
