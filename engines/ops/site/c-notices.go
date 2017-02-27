package site

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request) {
	session := sessions.GetSession(r)
	session.Set("hello", "world")
	p.Ctx.HTML(w, "site/notices/index", r.Context().Value(web.DATA))
}
