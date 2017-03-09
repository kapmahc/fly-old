package forum

import (
	"fmt"
	"net/http"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Server   *machinery.Server `inject:""`
	Cache    *web.Cache        `inject:""`
	I18n     *web.I18n         `inject:""`
	Settings *web.Settings     `inject:""`
	Db       *gorm.DB          `inject:""`
	Session  *auth.Session     `inject:""`
	Ctx      *web.Context      `inject:""`
	Mux      *web.Mux          `inject:""`
	Dao      *auth.Dao         `inject:""`
}

// Do background jobs
func (p *Engine) Do() {}

// Atom rss.atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	var items []*atom.Entry

	var articles []Article
	if err := p.Db.
		Select([]string{"id", "title", "summary", "updated_at"}).
		Order("updated_at DESC").Limit(32).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	for _, a := range articles {
		items = append(items, &atom.Entry{
			Title:     a.Title,
			ID:        fmt.Sprintf("articles-%d", a.ID),
			Published: atom.Time(a.UpdatedAt),
			Summary:   &atom.Text{Body: a.Summary},
		})
	}
	return items, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	urls := []stm.URL{
		{"loc": p.Ctx.URLFor("forum.articles.index")},
		{"loc": p.Ctx.URLFor("forum.tags.index")},
		{"loc": p.Ctx.URLFor("forum.comments.index")},
	}

	var articles []Article
	if err := p.Db.Select([]string{"id"}).Find(&articles).Error; err != nil {
		return nil, err
	}
	for _, a := range articles {
		urls = append(urls, stm.URL{"loc": p.Ctx.URLFor("forum.articles.show", "id", a.ID)})
	}

	var tags []Tag
	if err := p.Db.Select([]string{"id"}).Find(&tags).Error; err != nil {
		return nil, err
	}
	for _, t := range tags {
		urls = append(urls, stm.URL{"loc": p.Ctx.URLFor("forum.tags.show", "id", t.ID)})
	}
	return urls, nil
}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// NavBar nav-bar
func (p *Engine) NavBar(r *http.Request) ([]web.Link, *web.Dropdown) {
	dsh := web.Dropdown{
		Label: "forum.dashboard.title",
		Items: []*web.Link{
			&web.Link{Label: "forum.articles.new.title", Href: "forum.articles.new"},
			&web.Link{Label: "forum.dashboard.articles.title", Href: "forum.dashboard.articles"},
			&web.Link{Label: "forum.dashboard.comments.title", Href: "forum.dashboard.comments"},
		},
	}
	if p.Session.CheckAdmin(nil, r, false) {
		dsh.Items = append(dsh.Items, &web.Link{Label: "forum.dashboard.tags.title", Href: "forum.dashboard.tags"})
	}
	return []web.Link{
		{Label: "forum.articles.index.title", Href: "forum.articles.index"},
		{Label: "forum.comments.index.title", Href: "forum.comments.index"},
	}, &dsh
}

func init() {
	web.Register(&Engine{})
}
