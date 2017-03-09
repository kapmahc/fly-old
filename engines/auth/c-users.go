package auth

import (
	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexUsers(c *gin.Context) {

	var users []User
	err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error
	web.JSON(c, users, err)
}
