package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

type fmSignUp struct {
	Name                 string `form:"fullName" validate:"required,max=255"`
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

func (p *Engine) signIn(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.sign-in.title")
	p.Render.HTML(w, "auth/users/sign-in", data)
}

func (p *Engine) confirm(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.confirm.title")
	p.Render.HTML(w, "auth/users/confirm", data)
}

func (p *Engine) unlock(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.unlock.title")
	p.Render.HTML(w, "auth/users/unlock", data)
}

func (p *Engine) forgotPassword(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.forgot-password.title")
	p.Render.HTML(w, "auth/users/forgot-password", data)
}

func (p *Engine) resetPassword(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.reset-password.title")
	p.Render.HTML(w, "auth/users/reset-password", data)
}
