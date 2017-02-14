package web

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine interface {
	Beans() ([]*inject.Object, error)
	Mount(*mux.Router)
	Do()
	Shell() []cli.Command
	Atom() ([]*atom.Entry, error)
	Sitemap() ([]stm.URL, error)
}
