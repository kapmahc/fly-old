package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminPages(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.pages.index.title")
	tpl := "site-admin-pages-index"
	var pages []web.Page
	if err := p.Db.Order("loc ASC, sort_order DESC").Find(&pages).Error; err != nil {
		return tpl, err
	}
	data["items"] = pages
	return tpl, nil
}
