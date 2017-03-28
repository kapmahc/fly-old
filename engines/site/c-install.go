package site

import (
	"time"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmInstall struct {
	Title                string `form:"title" binding:"required"`
	SubTitle             string `form:"subTitle" binding:"required"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postInstall(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lang, "errors.forbidden")
	}

	var fm fmInstall
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	p.I18n.Set(lang, "site.title", fm.Title)
	p.I18n.Set(lang, "site.subTitle", fm.SubTitle)
	user, err := p.Dao.AddEmailUser("root", fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}
	for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
		role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
		if er != nil {
			return nil, er
		}
		if err = p.Dao.Allow(role.ID, user.ID, 50, 0, 0); err != nil {
			return nil, err
		}
	}
	if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil

}
