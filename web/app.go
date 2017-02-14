package web

import (
	"os"

	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

// New new application
func New() *Application {
	var g inject.Graph
	return &Application{
		router:  mux.NewRouter(),
		engines: make([]Engine, 0),
		graph:   &g,
	}
}

// Application application
type Application struct {
	router  *mux.Router
	engines []Engine
	graph   *inject.Graph
}

// Register register engine
func (p *Application) Register(ens ...Engine) {
	p.engines = append(p.engines, ens...)
	// for _, e := range ens {
	// 	e.Mount(p.router)
	// 	// -----------
	//
	// 	// -----------
	// 	args, err := eng.Beans()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	args = append(args, &inject.Object{Value: eng})
	// 	if err := p.graph.Provide(args...); err != nil {
	// 		return err
	// 	}
	// }
	// return nil
}

// Main main entry
func (p *Application) Main(usage, version string) error {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = version
	app.Usage = usage
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range p.engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
}
