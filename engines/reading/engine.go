package reading

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/unrolled/render"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Cache   *web.Cache     `inject:""`
	I18n    *web.I18n      `inject:""`
	Db      *gorm.DB       `inject:""`
	Session *auth.Session  `inject:""`
	Render  *render.Render `inject:""`
	Router  *mux.Router    `inject:""`
	UF      *auth.UrlFor   `inject:""`
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
		{"loc": p.UF.Path("reading.books.index")},
	}
	for _, b := range books {
		log.Println(p.UF.Path("reading.book.show", "id", b.ID))
		urls = append(
			urls,
			stm.URL{"loc": p.UF.Path("reading.book.show", "id", b.ID)},
		)
	}
	return urls, nil
}

func init() {
	web.Register(&Engine{})
}
