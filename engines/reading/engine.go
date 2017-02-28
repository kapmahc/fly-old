package reading

import (
	"net/http"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Cache   *web.Cache    `inject:""`
	I18n    *web.I18n     `inject:""`
	Db      *gorm.DB      `inject:""`
	Session *auth.Session `inject:""`
	Ctx     *web.Context  `inject:""`
	Mux     *web.Mux      `inject:""`
}

// Do background jobs
func (p *Engine) Do() {}

// Atom rss.atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	var books []Book
	if err := p.Db.Select([]string{"id"}).Find(&books).Error; err != nil {
		return nil, err
	}
	urls := []stm.URL{
		{"loc": p.Ctx.URLFor("reading.books.index")},
	}
	for _, b := range books {
		urls = append(
			urls,
			stm.URL{"loc": p.Ctx.URLFor("reading.book.show", "id", b.ID)},
		)
	}
	return urls, nil
}

// NavBar nav-bar
func (p *Engine) NavBar(r *http.Request) ([]web.Link, *web.Dropdown) {
	return []web.Link{
		{Label: "reading.books.index.title", Href: "reading.books.index"},
	}, nil
}

func init() {
	web.Register(&Engine{})
}
