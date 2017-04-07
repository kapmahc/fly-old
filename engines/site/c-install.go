package site

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/sky"
)

func (p *Engine) getInstall(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)
	fm := sky.NewForm(c)
	fm.Title(p.I18n.T(lang, "site.install.title"))
	data["form"] = fm
	c.HTML(http.StatusOK, "site/install", data)
	return nil
}

func (p *Engine) postInstall(c *sky.Context) error {
	return nil
}

func (p *Engine) mustDatabaseEmpty(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(lang, "errors.forbidden")
	}
	return c.Next()
}
