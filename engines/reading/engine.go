package reading

import (
	"net/http"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Cache   *web.Cache     `inject:""`
	I18n    *web.I18n      `inject:""`
	Db      *gorm.DB       `inject:""`
	Session *auth.Session  `inject:""`
	Render  *render.Render `inject:""`
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
	return p.indexBooks
}

func init() {
	web.Register(&Engine{})
}
