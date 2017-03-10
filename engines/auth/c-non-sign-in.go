package auth

import (
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

func (p *Engine) postUsersSignUp(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmSignUp
	var count int
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.
			Model(&User{}).
			Where("email = ?", fm.Email).
			Count(&count).Error
	}
	if err == nil && count > 0 {
		err = p.I18n.E(lang, "auth.errors.email-already-exists")
	}
	if err == nil {
		var user *User
		user, err = p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
		if err == nil {
			p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-up"))
			p.sendEmail(lang, user, actConfirm)
		}
	}
	web.TEXT(c, p.I18n.T(lang, "auth.messages.email-for-confirm"), err)
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Engine) postUsersSignIn(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmSignIn
	var user *User
	err := c.Bind(&fm)
	ip := c.ClientIP()
	if err == nil {
		user, err = p.Dao.GetByEmail(fm.Email)
		if err == nil {
			if !p.Security.Chk([]byte(fm.Password), user.Password) {
				p.Dao.Log(user.ID, ip, p.I18n.T(lang, "auth.logs.sign-in-failed"))
				err = p.I18n.E(lang, "auth.errors.email-password-not-match")
			}
		}
		if err == nil {
			if !user.IsConfirm() {
				err = p.I18n.E(lang, "auth.errors.user-not-confirm")
			}
		}
		if err == nil {
			if user.IsLock() {
				err = p.I18n.E(lang, "auth.errors.user-is-lock")
			}
		}
	}

	if err == nil {
		p.Dao.Log(user.ID, ip, p.I18n.T(lang, "auth.logs.sign-in-success"))
		user.SignInCount++
		user.LastSignInAt = user.CurrentSignInAt
		user.LastSignInIP = user.CurrentSignInIP
		now := time.Now()
		user.CurrentSignInAt = &now
		user.CurrentSignInIP = ip
		p.Db.Model(user).Updates(map[string]interface{}{
			"last_sign_in_at":    user.LastSignInAt,
			"last_sign_in_ip":    user.LastSignInIP,
			"current_sign_in_at": user.CurrentSignInAt,
			"current_sign_in_ip": user.CurrentSignInIP,
			"sign_in_count":      user.SignInCount,
		})

	}

	var tkn []byte
	if err == nil {
		cm := jws.Claims{}
		cm.Set(UID, user.UID)
		cm.Set(IsAdmin, p.Dao.Is(user.ID, RoleAdmin))
		tkn, err = p.Jwt.Sum(cm, time.Hour*24*7)
	}
	web.TEXT(c, string(tkn), err)
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Engine) getUsersConfirm(c *gin.Context) {
	token := c.Param("token")
	lang := c.MustGet(web.LOCALE).(string)
	user, err := p.parseToken(lang, token, actConfirm)
	if err == nil {
		if user.IsConfirm() {
			err = p.I18n.E(lang, "auth.errors.user-already-confirm")
		}
	}
	if err == nil {
		p.Db.Model(user).Update("confirmed_at", time.Now())
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.confirm"))
	}

	web.TEXT(c, p.I18n.T(lang, "auth.messages.confirm-success"), err)
}

func (p *Engine) postUsersConfirm(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmEmail
	var user *User
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetByEmail(fm.Email)
	}
	if err == nil {
		if user.IsConfirm() {
			err = p.I18n.E(lang, "auth.errors.user-already-confirm")
		}
	}
	if err == nil {
		p.sendEmail(lang, user, actConfirm)
	}
	web.TEXT(c, p.I18n.T(lang, "auth.messages.email-for-confirm"), err)
}

func (p *Engine) getUsersUnlock(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(lang, token, actUnlock)
	if err == nil {
		if !user.IsLock() {
			err = p.I18n.E(lang, "auth.errors.user-not-lock")
		}
	}
	if err == nil {
		p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.unlock"))
	}
	web.TEXT(c, p.I18n.T(lang, "auth.messages.unlock-success"), err)
}

func (p *Engine) postUsersUnlock(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmEmail
	var user *User
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetByEmail(fm.Email)
	}
	if err == nil {
		if !user.IsLock() {
			err = p.I18n.E(lang, "auth.errors.user-not-lock")
		}
	}
	if err == nil {
		p.sendEmail(lang, user, actUnlock)

	}
	web.TEXT(c, p.I18n.T(lang, "auth.messages.email-for-unlock"), err)
}

func (p *Engine) postUsersForgotPassword(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmEmail
	var user *User
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.Dao.GetByEmail(fm.Email)
	}
	if err == nil {
		p.sendEmail(lang, user, actResetPassword)
	}
	web.TEXT(c, p.I18n.T(lang, "auth.messages.email-for-reset-password"), err)
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersResetPassword(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmResetPassword
	var user *User
	err := c.Bind(&fm)
	if err == nil {
		user, err = p.parseToken(lang, fm.Token, actResetPassword)
	}
	if err == nil {
		p.Db.Model(user).Update("password", p.Security.Sum([]byte(fm.Password)))
		p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.reset-password"))
	}

	web.TEXT(c, p.I18n.T(lang, "auth.messages.reset-password-success"), err)
}
