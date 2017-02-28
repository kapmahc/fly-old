package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
)

func (p *Engine) dashboard(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	p.Ctx.HTML(w, "site/dashboard", r.Context().Value(web.DATA))
}

func (p *Engine) home(w http.ResponseWriter, r *http.Request) {

	rt := p.Mux.Router.Get(viper.GetString("server.home"))
	if rt == nil {
		p.Ctx.NotFound(w)
		return
	}
	rt.GetHandler().ServeHTTP(w, r)
}
