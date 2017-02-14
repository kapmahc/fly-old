package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) getSignIn(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA)
	p.R.HTML(w, http.StatusOK, "auth/users/sign-in", data)
}
