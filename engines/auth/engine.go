package auth

import (
	"net/http"
	"reflect"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Server   *machinery.Server `inject:""`
	Cache    *web.Cache        `inject:""`
	I18n     *web.I18n         `inject:""`
	Settings *web.Settings     `inject:""`
	Db       *gorm.DB          `inject:""`
	Session  *Session          `inject:""`
	Render   *render.Render    `inject:""`
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
	return p.indexNotices
}

func init() {
	en := Engine{}
	web.Register(&en)
	viper.SetDefault("server.home", reflect.ValueOf(&en).Elem().Type().PkgPath())
}
