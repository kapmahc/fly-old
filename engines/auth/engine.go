package auth

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
}

// Beans beans
func (p *Engine) Beans() ([]*inject.Object, error) {
	return []*inject.Object{}, nil
}

// Mount mount
func (p *Engine) Mount(*mux.Router) {}

// Do worker
func (p *Engine) Do() {}

// Shell shell
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Atom rss.atom
func (p *Engine) Atom() ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}
