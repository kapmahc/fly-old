package site

import (
	"github.com/kapmahc/fly/engines/auth"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminUsers(c *gin.Context) (interface{}, error) {
	var items []auth.User
	err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error
	return items, err
}
