package shop

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
	Db   *gorm.DB  `inject:""`
	I18n *web.I18n `inject:""`
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
	return []stm.URL{}, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if _, ok := c.Get(auth.CurrentUser); !ok {
		return nil
	}
	dd := web.Dropdown{
		Label: "shop.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/shop/addresses", Label: "shop.addresses.index.title"},
			&web.Link{Href: "/shop/orders", Label: "shop.orders.index.title"},
			&web.Link{Href: "/shop/return-authorizations", Label: "shop.return-authorizations.index.title"},
			nil,
			&web.Link{Href: "/shop/return-authorizations/new", Label: "shop.return-authorizations.new.title"},
		},
	}
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {

	}
	return &dd
}

func init() {
	web.Register(&Engine{})
}
