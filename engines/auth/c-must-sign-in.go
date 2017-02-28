package auth

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) signOut(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	ss := sessions.GetSession(r)
	ss.Clear()

	user := p.Session.CurrentUser(r)
	lng := r.Context().Value(web.LOCALE).(string)
	p.Dao.Log(user.ID, p.Ctx.ClientIP(r), p.I18n.T(lng, "auth.logs.sign-out"))
	p.Ctx.JSON(w, web.H{"ok": true})
}

type fmInfo struct {
	Name string `form:"name" validate:"required,max=255"`
	Home string `form:"home" validate:"max=255"`
	Logo string `form:"logo" validate:"max=255"`
}

func (p *Engine) info(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.info.title")

	if r.Method == http.MethodPost {
		var fm fmInfo
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			user := p.Session.CurrentUser(r)
			err = p.Db.Model(user).Where("id = ?", user.ID).Updates(map[string]interface{}{
				"home": fm.Home,
				"logo": fm.Logo,
				"name": fm.Name,
			}).Error
		}
		if err != nil {
			data[web.ERROR] = err.Error()
		}
	}

	p.Ctx.HTML(w, "auth/users/info", data)
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" validate:"required"`
	NewPassword          string `form:"newPassword" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=NewPassword"`
}

func (p *Engine) changePassword(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.change-password.title")

	if r.Method == http.MethodPost {
		var fm fmChangePassword
		err := p.Ctx.Bind(&fm, r)
		user := p.Session.CurrentUser(r)
		if err == nil {
			if !p.Security.Chk([]byte(fm.CurrentPassword), user.Password) {
				err = p.I18n.E(lang, "auth.errors.bad-password")
			}
		}
		if err == nil {
			err = p.Db.Model(user).Where("id = ?", user.ID).Update("password", p.Security.Sum([]byte(fm.NewPassword))).Error
		}

		if err == nil {
			data[web.WARNING] = p.I18n.T(lang, "auth.messages.change-password-success")
		} else {
			data[web.ERROR] = err.Error()
		}
	}

	p.Ctx.HTML(w, "auth/users/change-password", data)
}

func (p *Engine) logs(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	var logs []Log
	user := p.Session.CurrentUser(r)
	err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").
		Find(&logs).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["logs"] = logs
	data["title"] = p.I18n.T(lang, "auth.users.logs.title")
	p.Ctx.HTML(w, "auth/users/logs", data)
}
