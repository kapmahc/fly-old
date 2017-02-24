package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) install(w http.ResponseWriter, r *http.Request) {
	var ct int
	err := p.Db.Model(&User{}).Count(&ct).Error
	if !p.Render.Check(w, err) {
		return
	}

	if ct > 0 {
		p.Render.NotFound(w)
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "site.install.title")
	p.Render.HTML(w, "auth/site/install", data)
}
