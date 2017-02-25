package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) signUp(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "auth.users.sign-up.title")
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
