package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/widgets"
)

type fmSignIn struct {
	Email      string `form:"email" validate:"required,email"`
	Password   string `form:"password" validate:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Engine) getUsersSignIn(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	title := p.I18n.T(lang, "auth.users.sign-in.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("auth.users.sign-in"),
		p.Layout.URLFor("site.dashboard"),
		title,
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), ""),
	)
	c.HTML(http.StatusOK, "auth/users/non-sign-in", data)
	return nil
}
func (p *Engine) postUsersSignIn(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return err
	}

	user, err := p.Dao.SignIn(lang, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return err
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     TOKEN,
		Value:    string(tkn),
		Path:     "/",
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, sky.H{})
	return nil
}

type fmSignUp struct {
	Name                 string `form:"name" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) getUsersSignUp(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	title := p.I18n.T(lang, "auth.users.sign-up.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("auth.users.sign-up"),
		p.Layout.URLFor("auth.users.sign-in"),
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.fullName"), ""),
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	c.HTML(http.StatusOK, "auth/users/non-sign-in", data)
	return nil
}

func (p *Engine) postUsersSignUp(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return err
	}

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return p.I18n.E(lang, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-up"))
	p.sendEmail(lang, user, actConfirm)
	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "auth.messages.email-for-confirm")})
	return nil
}

type fmEmail struct {
	Email string `form:"email" validate:"required,email"`
}

func (p *Engine) getUsersEmailForm(act string) sky.Handler {
	return func(c *sky.Context) error {
		data := c.Get(sky.DATA).(sky.H)
		lang := c.Get(sky.LOCALE).(string)

		title := p.I18n.T(lang, "auth.users."+act+".title")
		data["title"] = title
		data["form"] = widgets.NewForm(
			c.Request,
			lang,
			p.Layout.URLFor("auth.users."+act),
			p.Layout.URLFor("auth.users.sign-in"),
			title,
			widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		)
		c.HTML(http.StatusOK, "auth/users/non-sign-in", data)
		return nil
	}
}

func (p *Engine) getUsersConfirm(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	user, err := p.parseToken(lang, c.Param("token"), actConfirm)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return p.I18n.E(lang, "auth.errors.user-already-confirm")
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.confirm"))

	c.Redirect(http.StatusFound, p.Layout.URLFor("auth.users.sign-in"))
	return nil
}

func (p *Engine) postUsersConfirm(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}

	if user.IsConfirm() {
		return p.I18n.E(lang, "auth.errors.user-already-confirm")
	}

	p.sendEmail(lang, user, actConfirm)
	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "auth.messages.email-for-confirm")})
	return nil
}

func (p *Engine) getUsersUnlock(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	user, err := p.parseToken(lang, c.Param("token"), actUnlock)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(lang, "auth.errors.user-not-lock")
	}

	p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.unlock"))

	c.Redirect(http.StatusFound, p.Layout.URLFor("auth.users.sign-in"))

	return nil
}

func (p *Engine) postUsersUnlock(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(lang, "auth.errors.user-not-lock")
	}
	p.sendEmail(lang, user, actUnlock)

	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "auth.messages.email-for-unlock")})
	return nil
}

func (p *Engine) postUsersForgotPassword(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmEmail
	var user *User
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	p.sendEmail(lang, user, actResetPassword)

	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "auth.messages.email-for-reset-password")})
	return nil
}

type fmResetPassword struct {
	Token                string `form:"token" validate:"required"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) getUsersResetPassword(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	title := p.I18n.T(lang, "auth.users.reset-password.title")
	data["title"] = title
	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("auth.users.reset-password"),
		p.Layout.URLFor("auth.users.sign-in"),
		title,
		widgets.NewHiddenField("token", c.Param("token")),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	c.HTML(http.StatusOK, "auth/users/non-sign-in", data)
	return nil
}

func (p *Engine) postUsersResetPassword(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.parseToken(lang, fm.Token, actResetPassword)
	if err != nil {
		return err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.reset-password"))

	c.JSON(http.StatusOK, sky.H{"message": p.I18n.T(lang, "auth.messages.reset-password-success")})
	return nil
}
