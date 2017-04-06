package site

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/cache"
	"github.com/kapmahc/sky/i18n"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/settings"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Db         *gorm.DB           `inject:""`
	Settings   *settings.Settings `inject:""`
	CacheStore cache.Store        `inject:""`
	Cache      *cache.Cache       `inject:""`
	I18n       *i18n.I18n         `inject:""`
	Queue      job.Queue          `inject:""`
	Server     *job.Server        `inject:""`
	Jwt        *auth.Jwt          `inject:""`
	Matcher    language.Matcher   `inject:""`
	Router     *mux.Router        `inject:""`
	Layout     *auth.Layout       `inject:""`
}

// Map map object
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Workers job workers
func (p *Engine) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Application application
func (p *Engine) Application(*sky.Context) []*sky.Dropdown {
	return nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(*sky.Context) []*sky.Dropdown {
	return nil
}

func init() {
	sky.Register(&Engine{})
}
