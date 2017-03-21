package vpn

import (
	"net/http"
	"time"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexUsers(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.users.index.title")
	tpl := "ops-vpn-users-index"
	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
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
	Details              string `form:"details"`
	Enable               bool   `form:"enable"`
	StartUp              string `form:"startUp"`
	ShutDown             string `form:"shutDown"`
}

func (p *Engine) createUser(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "ops-vpn-users-new"

	now := time.Now()
	data["startUp"] = now.Format(web.FormatDateInput)
	data["shutDown"] = now.AddDate(1, 0, 0).Format(web.FormatDateInput)

	if c.Request.Method == http.MethodPost {
		var fm fmUserNew
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
		if err != nil {
			return tpl, err
		}
		shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
		if err != nil {
			return tpl, err
		}
		user := User{
			FullName: fm.FullName,
			Email:    fm.Email,
			Details:  fm.Details,
			Enable:   fm.Enable,
			StartUp:  startUp,
			ShutDown: shutDown,
		}
		if err := user.SetPassword(fm.Password); err != nil {
			return tpl, err
		}
		if err := p.Db.Create(&user).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/vpn/users")
		return "", nil
	}
	return tpl, nil
}

type fmUserEdit struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable"`
	StartUp  string `form:"startUp"`
	ShutDown string `form:"shutDown"`
}

func (p *Engine) updateUser(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.users.edit.title")
	tpl := "ops-vpn-users-edit"
	id := c.Param("id")

	var item User
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item
	data["startUp"] = item.StartUp.Format(web.FormatDateInput)
	data["shutDown"] = item.ShutDown.Format(web.FormatDateInput)

	if c.Request.Method == http.MethodPost {
		var fm fmUserEdit
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
		if err != nil {
			return tpl, err
		}
		shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
		if err != nil {
			return tpl, err
		}
		if err := p.Db.Model(&User{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"full_name": fm.FullName,
				"enable":    fm.Enable,
				"start_up":  startUp,
				"shut_down": shutDown,
				"details":   fm.Details,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/vpn/users")
		return "", nil
	}

	return tpl, nil
}

type fmUserResetPassword struct {
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) resetUserPassword(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.users.reset-password.title")
	tpl := "ops-vpn-users-reset-password"
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
		c.Redirect(http.StatusFound, "/ops/vpn/users")
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
	data["title"] = p.I18n.T(lang, "ops.vpn.users.change-password.title")
	tpl := "ops-vpn-users-change-password"

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
			return tpl, p.I18n.E(lang, "ops.vpn.users.email-password-not-match")
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
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(User{}).Error
	return gin.H{}, err
}
