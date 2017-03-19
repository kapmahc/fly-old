package forum

import (
	"fmt"

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
		{"loc": "/forum/articles"},
		{"loc": "/forum/tags"},
		{"loc": "/forum/comments"},
	}

	var articles []Article
	if err := p.Db.Select([]string{"id"}).Find(&articles).Error; err != nil {
		return nil, err
	}
	for _, a := range articles {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/articles/show/%d", a.ID)})
	}

	var tags []Tag
	if err := p.Db.Select([]string{"id"}).Find(&tags).Error; err != nil {
		return nil, err
	}
	for _, t := range tags {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/tags/show/%d", t.ID)})
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *gin.Context) *web.Dropdown {
	if _, ok := c.Get(auth.CurrentUser); !ok {
		return nil
	}
	dd := web.Dropdown{
		Label: "forum.dashboard.title",
		Links: []*web.Link{
			&web.Link{Href: "/forum/articles/new", Label: "forum.articles.new.title"},
			nil,
			&web.Link{Href: "/forum/articles/my", Label: "forum.articles.my.title"},
			&web.Link{Href: "/forum/comments/my", Label: "forum.comments.my.title"},
		},
	}
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		dd.Links = append(
			dd.Links,
			nil,
			&web.Link{Href: "/forum/tags?act=manage", Label: "forum.tags.index.title"},
		)
	}
	return &dd
}

func init() {
	web.Register(&Engine{})
}
