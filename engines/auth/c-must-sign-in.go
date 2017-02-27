package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) signOut(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
}

func (p *Engine) profile(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
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
	if !p.Render.Check(w, err) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["logs"] = logs
	data["title"] = p.I18n.T(lang, "auth.users.logs.title")
	p.Render.HTML(w, "auth/users/logs", data)
}
