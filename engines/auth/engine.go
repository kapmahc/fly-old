package auth

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Server    *machinery.Server `inject:""`
	Cache     *web.Cache        `inject:""`
	I18n      *web.I18n         `inject:""`
	Settings  *web.Settings     `inject:""`
	Db        *gorm.DB          `inject:""`
	Session   *Session          `inject:""`
	Render    *web.Render       `inject:""`
	Router    *mux.Router       `inject:""`
	Dao       *Dao              `inject:""`
	UF        *web.UrlFor       `inject:""`
	Validator *web.Validator    `inject:""`
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

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
