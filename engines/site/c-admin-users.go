package site

import (
	"github.com/kapmahc/fly/engines/auth"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminUsers(c *gin.Context, lang string, data gin.H) (string, error) {
	var items []auth.User
	err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error
	data["users"] = items
	data["title"] = p.I18n.T(lang, "site.admin.users.index.title")
	return "site-admin-users-index", err
}
