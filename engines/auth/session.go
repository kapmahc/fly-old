package auth

import (
	"context"
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/fly/web"
)

const (
	// UID uid
	UID         = "uid"
	currentUser = "currentUser"
	isAdmin     = "is_admin"
)

// Session session
type Session struct {
	Dao  *Dao         `inject:""`
	Ctx  *web.Context `inject:""`
	I18n *web.I18n    `inject:""`
}

// CurrentUser current-user
func (p *Session) CurrentUser(r *http.Request) *User {
	u := r.Context().Value(web.DATA).(web.H)[currentUser]
	if u == nil {
		return nil
	}
	return u.(*User)
}

// CheckSignIn is sign-in?
func (p *Session) CheckSignIn(w http.ResponseWriter, r *http.Request, must bool) bool {
	u := p.CurrentUser(r)
	if u != nil {
		return true
	}
	if must {
		lang := r.Context().Value(web.LOCALE).(string)
		p.Ctx.Render.Text(w, http.StatusForbidden, p.I18n.T(lang, "auth.errors.please-sign-in"))
	}
	return false
}

// CheckAdmin is admin?
func (p *Session) CheckAdmin(w http.ResponseWriter, r *http.Request, must bool) bool {
	if p.CheckSignIn(w, r, must) {
		is := r.Context().Value(web.DATA).(web.H)[isAdmin]
		if is != nil && is.(bool) {
			return true
		}
		if must {
			lang := r.Context().Value(web.LOCALE).(string)
			p.Ctx.Render.Text(w, http.StatusForbidden, p.I18n.T(lang, "auth.errors.not-allow"))
		}
	}
	return false
}

// CurrentUserMiddleware current-user middleware
func (p *Session) CurrentUserMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ss := sessions.GetSession(r)
	uid := ss.Get(UID)
	if uid != nil {
		if user, err := p.Dao.GetUserByUID(uid.(string)); err == nil {
			data := r.Context().Value(web.DATA).(web.H)
			data[currentUser] = user
			data[isAdmin] = p.Dao.Is(user.ID, RoleAdmin)
			next(w, r.WithContext(context.WithValue(r.Context(), web.DATA, data)))
			return
		}
	}
	next(w, r)

}
