package vpn

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexUsers(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "vpn.users.index.title")
	tpl := "vpn-users-index"
	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmSignUp struct {
	FullName             string `form:"full_name" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}
