package auth

import (
	"net/http"

	"github.com/kapmahc/sky"
)

func (p *Engine) indexUsers(c *sky.Context) error {

	lang := c.Get(sky.LOCALE).(string)
	data := c.Get(sky.DATA).(sky.H)
	data["title"] = p.I18n.T(lang, "auth.users.index.title")

	var users []User
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		return err
	}
	data["items"] = users
	c.HTML(http.StatusOK, "auth/users/index", data)
	return nil
}
