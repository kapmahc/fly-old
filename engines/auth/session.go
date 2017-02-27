package auth

import (
	"context"
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
)

const (
	// UID uid
	UID = "uid"
	// CurrentUser current user
	CurrentUser = "currentUser"
	// IsAdmin is admin?
	IsAdmin = "is_admin"
)

type Session struct {
	Dao *Dao     `inject:""`
	Db  *gorm.DB `inject:""`
}

// Middleware current-user middleware
func (p *Session) Middleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ss := sessions.GetSession(r)
	if user, err := p.Dao.GetUserByUID(ss.Get(UID).(string)); err == nil {
		data := r.Context().Value(web.DATA).(web.H)
		data[CurrentUser] = user
		data[IsAdmin] = p.Dao.Is(user.ID, RoleAdmin)
		next(w, r.WithContext(context.WithValue(r.Context(), web.DATA, data)))
	} else {
		next(w, r)
	}

}
