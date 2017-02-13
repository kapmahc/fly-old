package shop

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
}

// Mount web mount points
func (p *Engine) Mount(*mux.Router) {

}

// Do background jobs
func (p *Engine) Do() {}

// Shell shell commands
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

// Home home
func (p *Engine) Home() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {}
}

// Dashboard dashboard
func (p *Engine) Dashboard(*http.Request) ([]web.Dropdown, error) {
	return []web.Dropdown{}, nil
}

func init() {
	web.Register(&Engine{})
}
