package auth

import (
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/fly/web"
)

type fmSignUp struct {
	Name                 string `form:"name" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) signUp(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.sign-up.title")

	if r.Method == http.MethodPost {
		var fm fmSignUp
		var count int
		err := p.Validator.Bind(&fm, r)
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
				p.Dao.Log(user.ID, p.Render.ClientIP(r), p.I18n.T(lang, "auth.logs.sign-up"))
				p.sendEmail(lang, user, actConfirm)
			}
		}

		if err == nil {
			data[web.INFO] = p.I18n.T(lang, "auth.messages.email-for-confirm")
		} else {
			data[web.ERROR] = err.Error()
		}
	}

	p.Render.HTML(w, "auth/users/sign-up", data)
}

type fmSignIn struct {
	Email      string `form:"email" validate:"required,email"`
	Password   string `form:"password" validate:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Engine) signIn(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.sign-in.title")

	if r.Method == http.MethodPost {
		var fm fmSignIn
		var user *User

		err := p.Validator.Bind(&fm, r)
		if err == nil {
			user, err = p.Dao.GetByEmail(fm.Email)
			ip := p.Render.ClientIP(r)
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
			ss := sessions.GetSession(r)
			ss.Set(UID, user.UID)
			p.Render.Redirect(w, r, "/")
			return
		} else {
			data[web.ERROR] = err.Error()
		}

	}

	p.Render.HTML(w, "auth/users/sign-in", data)
}

type fmEmail struct {
	Email string `form:"email" validate:"required,email"`
}

func (p *Engine) confirm(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.confirm.title")
	token := r.URL.Query().Get("token")
	if r.Method == http.MethodGet && token != "" {
		user, err := p.parseToken(lang, token, actConfirm)
		if err == nil {
			if user.IsConfirm() {
				err = p.I18n.E(lang, "auth.errors.user-already-confirm")
			}
		}
		if err == nil {
			p.Db.Model(user).Update("confirmed_at", time.Now())
			p.Dao.Log(user.ID, p.Render.ClientIP(r), p.I18n.T(lang, "auth.logs.confirm"))
			data[web.INFO] = p.I18n.T(lang, "auth.messages.confirm-success")
		} else {
			data[web.ERROR] = err.Error()
		}
	} else if r.Method == http.MethodPost {
		var fm fmEmail
		var user *User
		err := p.Validator.Bind(&fm, r)
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
			data[web.INFO] = p.I18n.T(lang, "auth.messages.email-for-confirm")
		} else {
			data[web.ERROR] = err.Error()
		}
	}
	p.Render.HTML(w, "auth/users/confirm", data)
}

func (p *Engine) unlock(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.unlock.title")
	token := r.URL.Query().Get("token")
	if r.Method == http.MethodGet && token != "" {
		user, err := p.parseToken(lang, token, actUnlock)
		if err == nil {
			if !user.IsLock() {
				err = p.I18n.E(lang, "auth.errors.user-not-lock")
			}
		}
		if err == nil {
			p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
			p.Dao.Log(user.ID, p.Render.ClientIP(r), p.I18n.T(lang, "auth.logs.unlock"))
			data[web.INFO] = p.I18n.T(lang, "auth.messages.unlock-success")
		} else {
			data[web.ERROR] = err.Error()
		}
	} else if r.Method == http.MethodPost {
		var fm fmEmail
		var user *User
		err := p.Validator.Bind(&fm, r)
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
			data[web.INFO] = p.I18n.T(lang, "auth.messages.email-for-unlock")
		} else {
			data[web.ERROR] = err.Error()
		}

	}
	p.Render.HTML(w, "auth/users/unlock", data)
}

func (p *Engine) forgotPassword(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.forgot-password.title")
	if r.Method == http.MethodPost {
		var fm fmEmail
		var user *User
		err := p.Validator.Bind(&fm, r)
		if err == nil {
			user, err = p.Dao.GetByEmail(fm.Email)
		}
		if err == nil {
			p.sendEmail(lang, user, actResetPassword)
			data[web.INFO] = p.I18n.T(lang, "auth.messages.email-for-reset-password")
		} else {
			data[web.ERROR] = err.Error()
		}
	}
	p.Render.HTML(w, "auth/users/forgot-password", data)
}

type fmResetPassword struct {
	Token                string `form:"token" validate:"required"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) resetPassword(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.reset-password.title")
	if r.Method == http.MethodGet {
		data["token"] = r.URL.Query().Get("token")
	} else {
		var fm fmResetPassword
		var user *User
		err := p.Validator.Bind(&fm, r)
		if err == nil {
			user, err = p.parseToken(lang, fm.Token, actResetPassword)
		}
		if err == nil {
			p.Db.Model(user).Update("password", p.Security.Sum([]byte(fm.Password)))
			p.Dao.Log(user.ID, p.Render.ClientIP(r), p.I18n.T(lang, "auth.logs.reset-password"))
			data[web.INFO] = p.I18n.T(lang, "auth.messages.reset-password-success")
		} else {
			data[web.ERROR] = err.Error()
		}
		data["token"] = fm.Token
	}

	p.Render.HTML(w, "auth/users/reset-password", data)
}
