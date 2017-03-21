package vpn

import (
	"github.com/SermoDigital/jose/crypto"
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
	Db       *gorm.DB             `inject:""`
	I18n     *web.I18n            `inject:""`
	Jwt      *auth.Jwt            `inject:""`
	Security *web.Security        `inject:""`
	Key      []byte               `inject:"jwt.key"`
	Method   crypto.SigningMethod `inject:"jwt.method"`
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
		{"loc": "/vpn/users/change-password"},
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		return &web.Dropdown{
			Label: "ops.vpn.dashboard.title",
			Links: []*web.Link{
				&web.Link{Href: "/vpn/users", Label: "ops.vpn.users.index.title"},
				&web.Link{Href: "/vpn/logs", Label: "ops.vpn.logs.index.title"},
				&web.Link{Href: "/vpn/readme", Label: "ops.vpn.readme.title"},
			},
		}
	}
	return nil
}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
