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
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/articles/%d", a.ID)})
	}

	var tags []Tag
	if err := p.Db.Select([]string{"id"}).Find(&tags).Error; err != nil {
		return nil, err
	}
	for _, t := range tags {
		urls = append(urls, stm.URL{"loc": fmt.Sprintf("/forum/tags/%d", t.ID)})
	}
	return urls, nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(*gin.Context) *web.Dropdown {
	return nil
}

func init() {
	web.Register(&Engine{})
}
