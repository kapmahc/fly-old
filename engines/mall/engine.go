package mall

import (
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

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
}

func init() {
	web.Register(&Engine{})
}
