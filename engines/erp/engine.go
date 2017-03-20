package erp

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
	return []stm.URL{}, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if admin, ok := c.Get(auth.IsAdmin); !ok || !admin.(bool) {
		return nil
	}
	return &web.Dropdown{
		Label: "erp.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/erp/shipping-methods", Label: "erp.shipping-methods.index.title"},
			&web.Link{Href: "/erp/payment-methods", Label: "erp.payment-methods.index.title"},
			&web.Link{Href: "/erp/zones", Label: "erp.zones.index.title"},
			nil,
			&web.Link{Href: "/erp/products", Label: "erp.products.index.title"},
			nil,
			&web.Link{Href: "/erp/orders", Label: "erp.orders.index.title"},
			&web.Link{Href: "/erp/return-authorizations", Label: "erp.return-authorizations.index.title"},
			nil,
			&web.Link{Href: "/erp/pos", Label: "erp.pos.index.title"},
		},
	}
}

func init() {
	web.Register(&Engine{})
}
