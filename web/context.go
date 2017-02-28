package web

import (
	"fmt"
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	validator "gopkg.in/go-playground/validator.v9"
)

// K context key type
type K string

const (
	// DATA data key
	DATA = K("data")
)

// Context http context helper
type Context struct {
	Router   *mux.Router         `inject:""`
	Render   *render.Render      `inject:""`
	I18n     *I18n               `inject:""`
	Validate *validator.Validate `inject:""`
	Decoder  *form.Decoder       `inject:""`
}

// URLFor url for
func (p *Context) URLFor(name string, args ...interface{}) string {
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

// Redirect redirect-to
func (p *Context) Redirect(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusFound)
}

// ClientIP client-ip
func (p *Context) ClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Error(err)
	}
	return ip
}

// Abort abort
func (p *Context) Abort(w http.ResponseWriter, lang, code string, args ...interface{}) {
	p.Render.Text(w, http.StatusInternalServerError, p.I18n.T(lang, code, args...))
}

// NotFound not-found
func (p *Context) NotFound(w http.ResponseWriter) {
	p.Render.Text(w, http.StatusNotFound, "not-found")
}

// HTML html
func (p *Context) HTML(w http.ResponseWriter, name string, data interface{}, options ...render.HTMLOptions) {
	p.Render.HTML(w, http.StatusOK, name, data, options...)
}

// JSON json
func (p *Context) JSON(w http.ResponseWriter, data interface{}) {
	p.Render.JSON(w, http.StatusOK, data)
}

// Check check error
func (p *Context) Check(w http.ResponseWriter, err error) bool {
	if err == nil {
		return true
	}
	log.Error(err)
	p.Render.Text(w, http.StatusInternalServerError, err.Error())
	return false
}

// Bind bind form
func (p *Context) Bind(fm interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	if err := p.Decoder.Decode(fm, r.Form); err != nil {
		return err
	}
	return p.Validate.Struct(fm)
}
