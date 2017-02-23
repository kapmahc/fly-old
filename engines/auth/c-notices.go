package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request) {
	p.Render.HTML(w, http.StatusOK, "auth/users/sign-up", r.Context().Value(web.DATA))
}
