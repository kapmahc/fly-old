package auth

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) signUp(w http.ResponseWriter, r *http.Request) {
	p.Render.HTML(w, http.StatusOK, "auth/users/sign-up", web.H{})
}

func (p *Engine) signIn(w http.ResponseWriter, r *http.Request) {
	p.Render.HTML(w, http.StatusOK, "auth/users/sign-in", web.H{})
}

func (p *Engine) confirm(http.ResponseWriter, *http.Request) {

}

func (p *Engine) unlock(http.ResponseWriter, *http.Request) {

}

func (p *Engine) forgotPassword(http.ResponseWriter, *http.Request) {

}

func (p *Engine) resetPassword(http.ResponseWriter, *http.Request) {

}
