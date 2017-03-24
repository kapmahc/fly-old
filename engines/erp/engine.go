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
	Db       *gorm.DB      `inject:""`
	I18n     *web.I18n     `inject:""`
	Jwt      *auth.Jwt     `inject:""`
	Dao      *auth.Dao     `inject:""`
	Settings *web.Settings `inject:""`
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
	if !p.isSeller(c) {
		return nil
	}
	// if IsPublic() || p.Dao.Is(user.(*auth.User), auth.RoleAdmin)
	dd := web.Dropdown{
		Label: "erp.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/erp/stores", Label: "erp.stores.index.title"},
			&web.Link{Href: "/erp/products", Label: "erp.products.index.title"},
			&web.Link{Href: "/erp/orders", Label: "erp.orders.index.title"},
			nil,
			&web.Link{Href: "/erp/pos", Label: "erp.pos.title"},
		},
	}

	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		dd.Links = append(
			dd.Links,
			nil,
			&web.Link{Href: "/erp/tags", Label: "erp.tags.index.title"},
		)
	}
	return &dd
}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
