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

func (p *Engine) info(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
}

func (p *Engine) changePassword(w http.ResponseWriter, r *http.Request) {
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
