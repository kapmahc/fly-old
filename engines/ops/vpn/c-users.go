package vpn

import (
	"time"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexUsers(c *gin.Context) (interface{}, error) {

	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
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

func (p *Engine) createUser(c *gin.Context) (interface{}, error) {

	var fm fmUserNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if err := p.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

type fmUserEdit struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable"`
	StartUp  string `form:"startUp"`
	ShutDown string `form:"shutDown"`
}

func (p *Engine) updateUser(c *gin.Context) (interface{}, error) {

	id := c.Param("id")

	var item User
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	var fm fmUserEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	startUp, err := time.Parse(web.FormatDateInput, fm.StartUp)
	if err != nil {
		return nil, err
	}
	shutDown, err := time.Parse(web.FormatDateInput, fm.ShutDown)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return gin.H{}, nil
}

type fmUserResetPassword struct {
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postResetUserPassword(c *gin.Context) (interface{}, error) {

	id := c.Param("id")

	var user User
	if err := p.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	var fm fmUserResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := user.SetPassword(fm.Password); err != nil {
		return nil, err
	}
	if err := p.Db.Model(&user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

type fmUserChangePassword struct {
	Email                string `form:"email" binding:"required,email"`
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) postChangeUserPassword(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmUserChangePassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if !user.ChkPassword(fm.CurrentPassword) {
		return nil, p.I18n.E(lang, "ops.vpn.users.email-password-not-match")
	}
	if err := user.SetPassword(fm.NewPassword); err != nil {
		return nil, err
	}

	if err := p.Db.Model(user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyUser(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(User{}).Error
	return gin.H{}, err
}
