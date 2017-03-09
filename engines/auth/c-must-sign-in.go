package auth

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) deleteUsersSignOut(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	lang := c.MustGet(web.LOCALE).(string)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-out"))
	web.JSON(c, nil, nil)
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	Home string `form:"home" binding:"max=255"`
	Logo string `form:"logo" binding:"max=255"`
}

func (p *Engine) postUsersInfo(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)

	var fm fmInfo
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Model(user).Where("id = ?", user.ID).Updates(map[string]interface{}{
			"home": fm.Home,
			"logo": fm.Logo,
			"name": fm.Name,
		}).Error
	}
	web.JSON(c, nil, err)
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) postUsersChangePassword(c *gin.Context) {

	user := c.MustGet(CurrentUser).(*User)
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmChangePassword
	err := c.Bind(&fm)
	if err == nil {
		if !p.Security.Chk([]byte(fm.CurrentPassword), user.Password) {
			err = p.I18n.E(lang, "auth.errors.bad-password")
		}
	}
	if err == nil {
		err = p.Db.Model(user).Where("id = ?", user.ID).Update("password", p.Security.Sum([]byte(fm.NewPassword))).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) getUsersLogs(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)

	var logs []Log
	err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").
		Find(&logs).Error
	web.JSON(c, logs, err)
}
