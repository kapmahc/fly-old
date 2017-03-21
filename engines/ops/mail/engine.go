package mail

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
	Db   *gorm.DB  `inject:""`
	I18n *web.I18n `inject:""`
	Jwt  *auth.Jwt `inject:""`
}

// RegisterWorker register worker
func (p *Engine) RegisterWorker() {

}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/ops/mail/users/change-password"},
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		return &web.Dropdown{
			Label: "ops.mail.dashboard.title",
			Links: []*web.Link{
				&web.Link{Href: "/ops/mail/domains", Label: "ops.mail.domains.index.title"},
				&web.Link{Href: "/ops/mail/users", Label: "ops.mail.users.index.title"},
				&web.Link{Href: "/ops/mail/aliases", Label: "ops.mail.aliases.index.title"},
				nil,
				&web.Link{Href: "/ops/mail/readme", Label: "ops.mail.readme.title"},
			},
		}
	}
	return nil
}

func init() {
	web.Register(&Engine{})
}
