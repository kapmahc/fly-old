package auth

import (
	"net/http"

	"github.com/kapmahc/fly-bak/web"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/widgets"
)

func (p *Engine) getUsersLogs(c *sky.Context) error {
	user := c.Get(CurrentUser).(*User)
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)
	data["title"] = p.I18n.T(lang, "auth.users.logs.title")
	var logs []Log
	if err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error; err != nil {
		return err
	}
	data["items"] = logs

	c.HTML(http.StatusOK, "auth/users/logs", data)
	return nil
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	Home string `form:"home" binding:"max=255"`
	Logo string `form:"logo" binding:"max=255"`
}

func (p *Engine) getUsersInfo(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)
	user := c.Get(CurrentUser).(*User)

	email := widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), user.Email)
	email.Readonly()
	title := p.I18n.T(lang, "auth.users.info.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("auth.users.info"),
		"",
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.fullName"), user.Name),
		email,
		widgets.NewTextField("home", p.I18n.T(lang, "auth.attributes.user.home"), user.Home),
		widgets.NewTextField("logo", p.I18n.T(lang, "auth.attributes.user.logo"), user.Logo),
	)
	c.HTML(http.StatusOK, "form", data)
	return nil
}
func (p *Engine) postUsersInfo(c *sky.Context) error {
	user := c.Get(CurrentUser).(*User)
	lang := c.Get(sky.LOCALE).(string)

	var fm fmInfo
	if err := c.Bind(&fm); err != nil {
		return err
	}

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		"home": fm.Home,
		"logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.update-info"))

	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "success")})
	return nil
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) getUsersChangePassword(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	title := p.I18n.T(lang, "auth.users.change-password.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("auth.users.change-password"),
		"",
		title,
		widgets.NewPasswordField("currentPassword", p.I18n.T(lang, "attributes.currentPassword"), ""),
		widgets.NewPasswordField("newPassword", p.I18n.T(lang, "attributes.newPassword"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	c.HTML(http.StatusOK, "form", data)
	return nil
}

func (p *Engine) postUsersChangePassword(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	user := c.Get(CurrentUser).(*User)
	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return p.I18n.E(lang, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "success")})
	return nil
}

func (p *Engine) deleteUsersSignOut(c *sky.Context) error {
	user := c.Get(CurrentUser).(*User)
	lang := c.Get(web.LOCALE).(string)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-out"))
	c.JSON(http.StatusOK, sky.H{})
	return nil
}
