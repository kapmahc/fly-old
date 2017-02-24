package site

import (
	"net/http"

	"github.com/spf13/viper"
)

func (p *Engine) home(w http.ResponseWriter, r *http.Request) {

	rt := p.Router.Get(viper.GetString("server.home"))
	if rt == nil {
		p.Render.NotFound(w)
		return
	}
	rt.GetHandler().ServeHTTP(w, r)
}
