package site

import (
	"net/http"
	"time"

	"github.com/kapmahc/fly/engines/auth"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmInstall struct {
	Title                string `form:"title" binding:"required"`
	SubTitle             string `form:"subTitle" binding:"required"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) formInstall(c *gin.Context, l string, d gin.H) (string, error) {
	tpl := "site-install"
	d["title"] = p.I18n.T(l, "site.install.title")
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return tpl, err
	}
	if count > 0 {
		return tpl, p.I18n.E(l, "errors.forbidden")
	}
	if c.Request.Method == http.MethodPost {
		var fm fmInstall
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		p.I18n.Set(l, "site.title", fm.Title)
		p.I18n.Set(l, "site.subTitle", fm.SubTitle)
		user, err := p.Dao.AddEmailUser("root", fm.Email, fm.Password)
		if err != nil {
			return tpl, err
		}
		for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
			role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
			if err == nil {
				er = p.Dao.Allow(role.ID, user.ID, 50, 0, 0)
			}
			if er != nil {
				return tpl, er
			}
		}
		if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/users/sign-in")
		return "", nil
	}
	return tpl, nil
}
