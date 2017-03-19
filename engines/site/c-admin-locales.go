package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminLocales(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.locales.index.title")
	var items []web.Locale
	err := p.Db.Select([]string{"id", "code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").Find(&items).Error
	data["items"] = items
	return "site-admin-locales-index", err
}

func (p *Engine) deleteAdminLocales(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(&web.Locale{}).Error

	return gin.H{}, err
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

func (p *Engine) formAdminLocales(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "site-admin-locales-edit"
	if c.Request.Method == http.MethodPost {
		var fm fmLocale
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		data["code"] = fm.Code
		if err := p.I18n.Set(lang, fm.Code, fm.Message); err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/locales")
		return "", nil
	}

	data["code"] = c.Request.URL.Query().Get("code")
	return tpl, nil
}
