package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexUsers(c *gin.Context) {

	var users []User
	err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error
	web.JSON(c, users, err)
}
