package auth

import (
	"net/http"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/fly/web"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Server *machinery.Server `inject:""`
	Cache  *web.Cache        `inject:""`
	I18n   *web.I18n         `inject:""`
	R      *render.Render    `inject:""`
	RT     *mux.Router       `inject:""`
}

// Do background jobs
func (p *Engine) Do() {}

// Atom rss.atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Home home
func (p *Engine) Home() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {}
}

// Dashboard dashboard
func (p *Engine) Dashboard(*http.Request) ([]web.Dropdown, error) {
	return []web.Dropdown{}, nil
}

func init() {
	web.Register(&Engine{})
}
