package site

import (
	"github.com/go-martini/martini"
	"github.com/kapmahc/fly/web"
)

// Engine auth engine
type Engine struct {
}

// Map objects
func (p *Engine) Map() interface{} {
	return func() {}
}

// Mount web-points
func (p *Engine) Mount(martini.Router) {}

func init() {
	web.Register(&Engine{})
}
