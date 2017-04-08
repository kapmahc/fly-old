package site

import (
	"net/http"

	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/widgets"
)

func (p *Engine) getAdminSiteInfo(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	title := p.I18n.T(lang, "site.admin.info.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("site.admin.info"),
		"",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "site.attributes.title"), p.I18n.T(lang, "site.title")),
		widgets.NewTextField("subTitle", p.I18n.T(lang, "site.attributes.subTitle"), p.I18n.T(lang, "site.subTitle")),
		widgets.NewTextField("keywords", p.I18n.T(lang, "site.attributes.keywords"), p.I18n.T(lang, "site.keywords")),
		widgets.NewTextarea("description", p.I18n.T(lang, "site.attributes.description"), p.I18n.T(lang, "site.description"), 3),
		widgets.NewTextField("copyright", p.I18n.T(lang, "site.attributes.copyright"), p.I18n.T(lang, "site.copyright")),
	)
	c.HTML(http.StatusOK, "form", data)
	return nil
}
func (p *Engine) postAdminSiteInfo(c *sky.Context) error {
	return nil
}
func (p *Engine) getAdminSiteAuthor(c *sky.Context) error {
	return nil
}
func (p *Engine) postAdminSiteAuthor(c *sky.Context) error {
	return nil
}
func (p *Engine) getAdminSiteSeo(c *sky.Context) error {
	return nil
}
func (p *Engine) postAdminSiteSeo(c *sky.Context) error {
	return nil
}
func (p *Engine) getAdminSiteSMTP(c *sky.Context) error {
	return nil
}
func (p *Engine) postAdminSiteSMTP(c *sky.Context) error {
	return nil
}
