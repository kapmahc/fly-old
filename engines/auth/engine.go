package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Dao      *Dao              `inject:""`
	Db       *gorm.DB          `inject:""`
	Security *web.Security     `inject:""`
	I18n     *web.I18n         `inject:""`
	Jwt      *Jwt              `inject:""`
	Server   *machinery.Server `inject:""`
	Uploader web.Uploader      `inject:""`
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/users"},
		{"loc": "/users/sign-in"},
		{"loc": "/users/sign-up"},
		{"loc": "/users/confirm"},
		{"loc": "/users/forgot-password"},
		{"loc": "/users/unlock"},
	}
	return urls, nil
}

func init() {
	web.Register(&Engine{})
}
