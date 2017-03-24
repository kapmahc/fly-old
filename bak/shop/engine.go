package shop

import (
	"fmt"

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
	var items []*atom.Entry

	var products []Product
	if err := p.Db.
		Select([]string{"id", "name", "description", "updated_at"}).
		Order("updated_at DESC").Limit(12).
		Find(&products).Error; err != nil {
		return nil, err
	}
	for _, v := range products {
		items = append(items, &atom.Entry{
			Title: v.Name,
			Link: []atom.Link{
				{Href: fmt.Sprintf("%s/shop/products/%d", web.Home(), v.ID)},
			},
			ID:        fmt.Sprintf("shop-products-%d", v.ID),
			Published: atom.Time(v.UpdatedAt),
			Summary:   &atom.Text{Body: v.Description},
		})
	}
	return items, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": "/shop/products"},
		{"loc": "/shop/vendors"},
		{"loc": "/shop/variants"},
		{"loc": "/shop/catalogs"},
	}
	// ----------
	var products []Product
	if err := p.Db.Select([]string{"id"}).Find(&products).Error; err != nil {
		return nil, err
	}
	for _, v := range products {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/shop/products/show/%d", v.ID)})
	}
	// -------------
	var catalogs []Catalog
	if err := p.Db.Select([]string{"id"}).Find(&catalogs).Error; err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/catalogs/show/%d", v.ID)})
	}
	// -----------
	var variants []Variant
	if err := p.Db.Select([]string{"id"}).Find(&variants).Error; err != nil {
		return nil, err
	}
	for _, v := range variants {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/variants/show/%d", v.ID)})
	}
	// -----------
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if _, ok := c.Get(auth.CurrentUser); !ok {
		return nil
	}
	return &web.Dropdown{
		Label: "shop.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/shop/addresses", Label: "shop.addresses.index.title"},
			&web.Link{Href: "/shop/orders", Label: "shop.orders.index.title"},
			&web.Link{Href: "/shop/return-authorizations", Label: "shop.return-authorizations.index.title"},
			nil,
			&web.Link{Href: "/shop/return-authorizations/new", Label: "shop.return-authorizations.new.title"},
		},
	}
}

func init() {
	web.Register(&Engine{})
}
