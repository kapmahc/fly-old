package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request) {
	p.Render.HTML(w, "auth/notices/index", r.Context().Value(web.DATA))
}
