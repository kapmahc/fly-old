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

func (p *Engine) postUsersSignUp(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSignUp
	var count int
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, p.I18n.E(lang, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-up"))
	p.sendEmail(lang, user, actConfirm)

	return gin.H{}, nil
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Engine) postUsersSignIn(c *gin.Context) (interface{}, error) {

	lang := c.MustGet(web.LOCALE).(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	user, err := p.Dao.SignIn(fm.Email, fm.Password, lang, c.ClientIP())
	if err != nil {
		return nil, err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"token": string(tkn),
		"user":  user,
	}, nil
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Engine) getUsersConfirm(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(lang, token, actConfirm)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if user.IsConfirm() {
		c.String(http.StatusInternalServerError, p.I18n.T(lang, "auth.errors.user-already-confirm"))
		return
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.confirm"))
	c.Redirect(http.StatusFound, "/")
}
func (p *Engine) postUsersConfirm(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}

	if user.IsConfirm() {
		return nil, p.I18n.E(lang, "auth.errors.user-already-confirm")
	}

	p.sendEmail(lang, user, actConfirm)

	return gin.H{}, nil
}

func (p *Engine) getUsersUnlock(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(lang, token, actUnlock)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !user.IsLock() {
		c.String(http.StatusInternalServerError, p.I18n.T(lang, "auth.errors.user-not-lock"))
		return
	}

	p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.unlock"))
	c.Redirect(http.StatusFound, "/")
}

func (p *Engine) postUsersUnlock(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(lang, "auth.errors.user-not-lock")
	}
	p.sendEmail(lang, user, actUnlock)

	return gin.H{}, nil
}

func (p *Engine) postUsersForgotPassword(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmEmail
	var user *User
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	p.sendEmail(lang, user, actResetPassword)

	return gin.H{}, nil
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersResetPassword(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	user, err := p.parseToken(lang, fm.Token, actResetPassword)
	if err != nil {
		return nil, err
	}
	p.Db.Model(user).Update("password", p.Security.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.reset-password"))

	return gin.H{}, nil
}
