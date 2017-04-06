package reading

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/job"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
}

// Map map object
func (p *Engine) Map(*inject.Graph) error {
	return nil
}

// Mount web mount points
func (p *Engine) Mount(*sky.Router) {

}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Workers job workers
func (p *Engine) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Navbar navbar
func (p *Engine) Navbar(*sky.Context) []*sky.Dropdown {
	return nil
}

// Dashboard dashboard
func (p *Engine) Dashboard(*sky.Context) []*sky.Dropdown {
	return nil
}

func init() {
	sky.Register(&Engine{})
}
