package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminLinks(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.links.index.title")
	tpl := "site-admin-links-index"
	var links []web.Link
	if err := p.Db.Order("loc ASC, sort_order DESC").Find(&links).Error; err != nil {
		return tpl, err
	}
	data["items"] = links
	return tpl, nil
}
