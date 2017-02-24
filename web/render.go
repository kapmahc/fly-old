package web

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/unrolled/render"
)

type Render struct {
	Render *render.Render `inject:""`
}

func (p *Render) NotFound(w http.ResponseWriter) {
	p.Render.Text(w, http.StatusNotFound, "not-found")
}

func (p *Render) HTML(w http.ResponseWriter, name string, data interface{}) {
	p.Render.HTML(w, http.StatusOK, name, data)
}

func (p *Render) Check(w http.ResponseWriter, err error) bool {
	if err == nil {
		return true
	}
	log.Error(err)
	p.Render.Text(w, http.StatusInternalServerError, err.Error())
	return false
}
