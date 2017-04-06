package site

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/i18n"
)

func (p *Engine) getInstall(c *sky.Context) error {
	c.HTML(http.StatusOK, "site/install", c.Get(sky.DATA))
	return nil
}

func (p *Engine) postInstall(c *sky.Context) error {
	return nil
}

func (p *Engine) mustDatabaseEmpty(c *sky.Context) error {
	lang := c.Get(i18n.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(lang, "errors.forbidden")
	}
	return c.Next()
}
