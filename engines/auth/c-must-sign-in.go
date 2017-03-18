package auth

import (
	"net/http"
	"time"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) deleteUsersSignOut(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	lang := c.MustGet(web.LOCALE).(string)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-out"))
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    TOKEN,
		Expires: time.Now().Add(time.Hour * -1),
		Path:    "/",
	})
	return gin.H{}, nil
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	Home string `form:"home" binding:"max=255"`
	Logo string `form:"logo" binding:"max=255"`
}

func (p *Engine) formUsersInfo(c *gin.Context, lang string, data gin.H) (string, error) {
	user := c.MustGet(CurrentUser).(*User)
	data["user"] = user
	data["title"] = p.I18n.T(lang, "auth.users.info.title")
	tpl := "auth-users-info"
	if c.Request.Method == http.MethodPost {

		var fm fmInfo
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(user).Updates(map[string]interface{}{
			"home": fm.Home,
			"logo": fm.Logo,
			"name": fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
	}
	return tpl, nil
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) formUsersChangePassword(c *gin.Context, lang string, data gin.H) (string, error) {
	user := c.MustGet(CurrentUser).(*User)
	data["title"] = p.I18n.T(lang, "auth.users.change-password.title")
	tpl := "auth-users-change-password"
	if c.Request.Method == http.MethodPost {
		var fm fmChangePassword
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if !p.Security.Chk([]byte(fm.CurrentPassword), user.Password) {
			return tpl, p.I18n.E(lang, "auth.errors.bad-password")
		}
		if err := p.Db.Model(user).
			Update("password", p.Security.Sum([]byte(fm.NewPassword))).Error; err != nil {
			return tpl, err
		}
		data[web.NOTICE] = p.I18n.T(lang, "success")
	}
	return tpl, nil
}

func (p *Engine) getUsersLogs(c *gin.Context, lang string, data gin.H) (string, error) {
	user := c.MustGet(CurrentUser).(*User)
	data["title"] = p.I18n.T(lang, "auth.users.logs.title")
	var logs []Log
	err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error
	data["logs"] = logs
	return "auth-users-logs", err
}
