package auth

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Dao  *Dao      `inject:""`
	I18n *web.I18n `inject:""`
	Db   *gorm.DB  `inject:""`
}

// Mount web mount points
func (p *Engine) Mount(*web.Router) {

}

// Workers job workers
func (p *Engine) Workers() map[string]web.Worker {
	return map[string]web.Worker{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Navbar navbar
func (p *Engine) Navbar(*web.Context) []*web.Dropdown {
	return []*web.Dropdown{}
}

// Dashboard dashboard
func (p *Engine) Dashboard(*web.Context) []*web.Dropdown {
	return []*web.Dropdown{}
}

func init() {
	web.Register(&Engine{})
}
