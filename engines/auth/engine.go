package auth

import (
	"net/http"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
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
	Session  *Session          `inject:""`
	Mux      *web.Mux          `inject:""`
	Dao      *Dao              `inject:""`
	Ctx      *web.Context      `inject:""`
	Security *web.Security     `inject:""`
	Jwt      *Jwt              `inject:""`
}

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
	return []web.Link{
			{Label: "auth.users.index.title", Href: "auth.users.index"},
		}, &web.Dropdown{
			Label: "auth.profile",
			Items: []*web.Link{
				&web.Link{Label: "auth.users.logs.title", Href: "auth.users.logs"},
				nil,
				&web.Link{Label: "auth.users.info.title", Href: "auth.users.info"},
				&web.Link{Label: "auth.users.change-password.title", Href: "auth.users.change-password"},
			},
		}
}

func init() {
	web.Register(&Engine{})
}
