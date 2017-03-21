package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmSignUp struct {
	Name                 string `form:"name" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) formUsersSignUp(c *gin.Context, lang string, data gin.H) (string, error) {
	tpl := "auth-users-sign-up"
	data["title"] = p.I18n.T(lang, "auth.users.sign-up.title")

	if c.Request.Method == http.MethodPost {
		var fm fmSignUp
		var count int
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if err := p.Db.
			Model(&User{}).
			Where("email = ?", fm.Email).
			Count(&count).Error; err != nil {
			return tpl, err
		}

		if count > 0 {
			return tpl, p.I18n.E(lang, "auth.errors.email-already-exists")
		}

		user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
		if err != nil {
			return tpl, err
		}

		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-up"))
		p.sendEmail(lang, user, actConfirm)
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.email-for-confirm")
	}
	return tpl, nil
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Engine) formUsersSignIn(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "auth.users.sign-in.title")
	tpl := "auth-users-sign-in"

	if c.Request.Method == http.MethodPost {
		var fm fmSignIn
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		user, err := p.Dao.SignIn(fm.Email, fm.Password, lang, c.ClientIP())
		if err != nil {
			return tpl, err
		}

		cm := jws.Claims{}
		cm.Set(UID, user.UID)
		// cm.Set("name", user.Name)
		// cm.Set(IsAdmin, p.Dao.Is(user.ID, RoleAdmin))
		tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
		if err != nil {
			return tpl, err
		}
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     TOKEN,
			Value:    string(tkn),
			Path:     "/",
			HttpOnly: true,
		})
		c.Redirect(http.StatusFound, "/dashboard")
		return "", nil
	}
	return tpl, nil
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Engine) formUsersConfirm(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "auth.users.confirm.title")
	data["action"] = "confirm"
	token := c.Param("token")
	method := c.Request.Method
	switch {
	case method == http.MethodGet && token != "":
		user, err := p.parseToken(lang, token, actConfirm)
		if err != nil {
			return tplUserEmail, err
		}
		if user.IsConfirm() {
			return tplUserEmail, p.I18n.E(lang, "auth.errors.user-already-confirm")
		}
		p.Db.Model(user).Update("confirmed_at", time.Now())
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.confirm"))
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.confirm-success")
	case method == http.MethodPost:
		var fm fmEmail
		if err := c.Bind(&fm); err != nil {
			return tplUserEmail, err
		}
		user, err := p.Dao.GetByEmail(fm.Email)
		if err != nil {
			return tplUserEmail, err
		}

		if user.IsConfirm() {
			return tplUserEmail, p.I18n.E(lang, "auth.errors.user-already-confirm")
		}

		p.sendEmail(lang, user, actConfirm)
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.email-for-confirm")
	}

	return tplUserEmail, nil
}

func (p *Engine) formUsersUnlock(c *gin.Context, lang string, data gin.H) (string, error) {
	data["action"] = "unlock"
	data["title"] = p.I18n.T(lang, "auth.users.unlock.title")
	method := c.Request.Method
	token := c.Param("token")
	switch {
	case method == http.MethodGet && token != "":
		user, err := p.parseToken(lang, token, actUnlock)
		if err != nil {
			return tplUserEmail, err
		}
		if !user.IsLock() {
			return tplUserEmail, p.I18n.E(lang, "auth.errors.user-not-lock")
		}

		p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.unlock"))
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.unlock-success")
	case method == http.MethodPost:
		var fm fmEmail
		if err := c.Bind(&fm); err != nil {
			return tplUserEmail, err
		}
		user, err := p.Dao.GetByEmail(fm.Email)
		if err != nil {
			return tplUserEmail, err
		}
		if !user.IsLock() {
			return tplUserEmail, p.I18n.E(lang, "auth.errors.user-not-lock")
		}
		p.sendEmail(lang, user, actUnlock)
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.email-for-unlock")
	}

	return tplUserEmail, nil
}

func (p *Engine) formUsersForgotPassword(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "auth.users.forgot-password.title")
	data["action"] = "forgot-password"
	if c.Request.Method == http.MethodPost {
		var fm fmEmail
		var user *User
		if err := c.Bind(&fm); err != nil {
			return tplUserEmail, err
		}
		user, err := p.Dao.GetByEmail(fm.Email)
		if err != nil {
			return tplUserEmail, err
		}
		p.sendEmail(lang, user, actResetPassword)
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.email-for-reset-password")
	}
	return tplUserEmail, nil
}

type fmResetPassword struct {
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) formUsersResetPassword(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "auth.users.reset-password.title")
	tpl := "auth-users-reset-password"
	if c.Request.Method == http.MethodPost {
		var fm fmResetPassword
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		user, err := p.parseToken(lang, c.Param("token"), actResetPassword)
		if err != nil {
			return tpl, err
		}
		p.Db.Model(user).Update("password", p.Security.Sum([]byte(fm.Password)))
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.reset-password"))
		data[web.NOTICE] = p.I18n.T(lang, "auth.messages.reset-password-success")
	}
	return tpl, nil
}

const (
	tplUserEmail = "auth-users-email-form"
)
