package blog

import (
	"net/http"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Server   *machinery.Server `inject:""`
	Cache    *web.Cache        `inject:""`
	I18n     *web.I18n         `inject:""`
	Settings *web.Settings     `inject:""`
	Db       *gorm.DB          `inject:""`
	Session  *auth.Session     `inject:""`
	Ctx      *web.Context      `inject:""`
	Mux      *web.Mux          `inject:""`
}

// Do background jobs
func (p *Engine) Do() {}

// Atom rss.atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	posts := p.getPosts()
	urls := []stm.URL{
		{"loc": p.Ctx.URLFor("blog.engine.home")},
	}
	for _, i := range posts {
		urls = append(
			urls,
			stm.URL{"loc": p.Ctx.URLFor("blog.show", "name", i.Href)},
		)
	}
	return urls, nil
}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// NavBar nav-bar
func (p *Engine) NavBar(r *http.Request) ([]web.Link, *web.Dropdown) {
	return []web.Link{
		{Label: "blog.index.title", Href: "blog.index"},
	}, nil
}

func init() {
	web.Register(&Engine{})
}
