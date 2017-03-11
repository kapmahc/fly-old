package blog

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	I18n *web.I18n `inject:""`
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
		{"loc": "/blogs"},
	}
	for _, lang := range viper.GetStringSlice("languages") {
		posts, err := p.getPosts(lang)
		if err != nil {
			return nil, err
		}
		for _, i := range posts {
			urls = append(
				urls,
				stm.URL{"loc": "/blog/" + i.Href},
			)
		}

	}
	return urls, nil
}

func init() {
	web.Register(&Engine{})
}
