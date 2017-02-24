package site

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

type installFm struct {
	Email                string `validate:"required,email"`
	Password             string `validate:"min=6,max=32"`
	PasswordConfirmation string `validate:"eqfield=Password"`
}

func (p *Engine) install(w http.ResponseWriter, r *http.Request) {
	var ct int
	err := p.Db.Model(&auth.User{}).Count(&ct).Error
	if !p.Render.Check(w, err) {
		return
	}

	if ct > 0 {
		p.Render.NotFound(w)
		return
	}

	if r.Method == http.MethodPost {
		var fm installFm
		p.Validator.Bind(fm, r)
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	data["title"] = p.I18n.T(lang, "site.install.title")
	p.Render.HTML(w, "site/install", data)
}
