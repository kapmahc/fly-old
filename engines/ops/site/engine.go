package site

import (
	"net/http"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/garyburd/redigo/redis"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
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
	Dao      *auth.Dao         `inject:""`
	Mux      *web.Mux          `inject:""`
	Redis    *redis.Pool       `inject:""`
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

// NavBar nav-bar
func (p *Engine) NavBar(r *http.Request) ([]web.Link, *web.Dropdown) {
	var dash *web.Dropdown
	if p.Session.CheckAdmin(nil, r, false) {
		dash = &web.Dropdown{
			Label: "site.dashboard.title",
			Items: []*web.Link{
				&web.Link{Label: "site.admin.status.title", Href: "site.admin.status"},
				nil,
				&web.Link{Label: "site.admin.info.title", Href: "site.admin.info"},
				&web.Link{Label: "site.admin.author.title", Href: "site.admin.author"},
				&web.Link{Label: "site.admin.seo.title", Href: "site.admin.seo"},
				&web.Link{Label: "site.admin.smtp.title", Href: "site.admin.smtp"},
				nil,
				&web.Link{Label: "site.admin.locales.index.title", Href: "site.admin.locales.index"},
				&web.Link{Label: "site.admin.users.index.title", Href: "site.admin.users.index"},
				nil,
				&web.Link{Label: "site.notices.index.title", Href: "site.notices.index"},
				&web.Link{Label: "site.leave-words.index.title", Href: "site.leave-words.index"},
			},
		}
	}
	return []web.Link{
		{Label: "site.notices.index.title", Href: "site.notices.index"},
	}, dash
}

func init() {
	web.Register(&Engine{})
}
