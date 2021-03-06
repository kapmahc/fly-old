package site

import (
	"github.com/garyburd/redigo/redis"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Cache    *web.Cache       `inject:""`
	I18n     *web.I18n        `inject:""`
	Settings *web.Settings    `inject:""`
	Db       *gorm.DB         `inject:""`
	Jwt      *auth.Jwt        `inject:""`
	Redis    *redis.Pool      `inject:""`
	Matcher  language.Matcher `inject:""`
	Dao      *auth.Dao        `inject:""`
	Queue    *web.Queue       `inject:""`
}

// RegisterWorker register worker
func (p *Engine) RegisterWorker() {

}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/"},
		{"loc": "/leave-words/new"},
		{"loc": "/notices"},
	}
	return urls, nil
}

func init() {
	web.Register(&Engine{})
}
