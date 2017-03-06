package auth

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) showUser(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)
	vars := mux.Vars(r)
	var user User
	if err := p.Db.
		Where("uid = ?", vars["uid"]).
		Find(&user).Error; err != nil {
		log.Error(err)
	}
	data["user"] = user
	data["title"] = user.Name
	p.Ctx.HTML(w, "auth/users/show", data)
}

func (p *Engine) indexUsers(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	var users []User
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		log.Error(err)
	}
	data["users"] = users
	data["title"] = p.I18n.T(lang, "auth.users.index.title")
	p.Ctx.HTML(w, "auth/users/index", data)
}
