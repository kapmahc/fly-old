package mail

import (
	"net/http"

	"github.com/kapmahc/fly/web"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexUsers(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.users.index.title")
	tpl := "ops-mail-users-index"
	var items []User
	if err := p.Db.Select([]string{"id", "name", "active"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmUserNew struct {
	FullName             string `form:"fullName" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
	Enable               bool   `form:"enable"`
	DomainID             uint   `form:"domainId"`
}

func (p *Engine) createUser(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "ops-mail-users-new"
	if c.Request.Method == http.MethodPost {
		var fm fmUserNew
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		user := User{
			FullName: fm.FullName,
			Email:    fm.Email,
			Enable:   fm.Enable,
			DomainID: fm.DomainID,
		}
		if err := user.SetPassword(fm.Password); err != nil {
			return tpl, err
		}
		if err := p.Db.Create(&user).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/users")
		return "", nil
	}
	return tpl, nil
}

type fmUserEdit struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Enable   bool   `form:"enable"`
}

func (p *Engine) updateUser(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.users.edit.title")
	tpl := "ops-mail-users-edit"
	id := c.Param("id")

	var item User
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmUserEdit
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&User{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"enable":    fm.Enable,
				"full_name": fm.FullName,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/users")
		return "", nil
	}

	return tpl, nil
}

type fmUserResetPassword struct {
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) resetUserPassword(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.users.reset-password.title")
	tpl := "ops-mail-users-reset-password"
	id := c.Param("id")

	var item User
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmUserResetPassword
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := item.SetPassword(fm.Password); err != nil {
			return tpl, err
		}
		if err := p.Db.Model(&item).
			Updates(map[string]interface{}{
				"password": item.Password,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/users")
		return "", nil
	}

	return tpl, nil
}

type fmUserChangePassword struct {
	Email                string `form:"email" binding:"required,email"`
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) changeUserPassword(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.users.change-password.title")
	tpl := "ops-mail-users-change-password"

	if c.Request.Method == http.MethodPost {
		var fm fmUserChangePassword
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		var user User
		if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
			return tpl, err
		}
		if !user.ChkPassword(fm.CurrentPassword) {
			return tpl, p.I18n.E(lang, "ops.mail.users.email-password-not-match")
		}
		if err := user.SetPassword(fm.NewPassword); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(user).
			Updates(map[string]interface{}{
				"password": user.Password,
			}).Error; err != nil {
			return tpl, err
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}

	return tpl, nil
}

func (p *Engine) destroyUser(c *gin.Context) (interface{}, error) {
	var user User
	if err := p.Db.
		Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return nil, err
	}
	var count int
	if err := p.Db.Model(&Alias{}).Where("destination = ?", user.Email).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(c.MustGet(web.LOCALE).(string), "errors.in-use")
	}
	err := p.Db.Delete(&user).Error
	return gin.H{}, err
}
