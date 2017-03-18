package reading

import (
	"fmt"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Engine engine
type Engine struct {
	Db *gorm.DB `inject:""`
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
	var books []Book
	if err := p.Db.Select([]string{"id"}).Find(&books).Error; err != nil {
		return nil, err
	}
	urls := []stm.URL{
		{"loc": "/reading/books"},
	}
	for _, b := range books {
		urls = append(
			urls,
			stm.URL{"loc": fmt.Sprintf("/reading/books/%d", b.ID)},
		)
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
