package web

import (
	"fmt"
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// UrlFor url helper
type UrlFor struct {
	Router *mux.Router `inject:""`
}

// Path path build
func (p *UrlFor) Path(name string, args ...interface{}) string {
	var pairs []string
	for _, arg := range args {
		pairs = append(pairs, fmt.Sprintf("%v", arg))
	}
	rt := p.Router.Get(name)
	if rt == nil {
		return "not-found"
	}
	url, err := rt.URL(pairs...)
	if err != nil {
		return err.Error()
	}
	return url.String()
}

type Render struct {
	Render *render.Render `inject:""`
	I18n   *I18n          `inject:""`
}

func (p *Render) ClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Error(err)
	}
	return ip
}

func (p *Render) Abort(w http.ResponseWriter, lang, code string, args ...interface{}) {
	p.Render.Text(w, http.StatusInternalServerError, p.I18n.T(lang, code, args...))
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
