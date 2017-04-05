package web

import "github.com/urfave/cli"

// Engine engine
type Engine interface {
	Shell() []cli.Command
}

// -----------------------------------------------------------------------------

var engines []Engine

// Register register engines
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

// Walk walk engines
func Walk(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
