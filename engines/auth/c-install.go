package auth

import (
	"net/http"
)

func (p *Engine) install(w http.ResponseWriter, r *http.Request) {
	var ct int
	err := p.Db.Model(&User{}).Count(&ct).Error
	if !p.Render.Check(w, err) {
		return
	}
}
