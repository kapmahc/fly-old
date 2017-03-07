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
	urls := []stm.URL{
		{"loc": p.Ctx.URLFor("auth.users.index")},
		{"loc": p.Ctx.URLFor("auth.users.sign-in")},
		{"loc": p.Ctx.URLFor("auth.users.sign-up")},
		{"loc": p.Ctx.URLFor("auth.users.confirm")},
		{"loc": p.Ctx.URLFor("auth.users.forgot-password")},
		{"loc": p.Ctx.URLFor("auth.users.unlock")},
	}
	var users []User
	if err := p.Db.Select([]string{"uid"}).Find(&users).Error; err != nil {
		return nil, err
	}
	for _, u := range users {
		urls = append(urls, stm.URL{"loc": p.Ctx.URLFor("auth.users.show", "uid", u.UID)})
	}
	return urls, nil
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
