package erp

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/urfave/cli"
)

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:  "erp",
			Usage: "erp operations",
			Subcommands: []cli.Command{
				{
					Name:    "seed",
					Usage:   "loads the seed data",
					Aliases: []string{"s"},
					Action:  auth.Action(p.loadSeed),
				},
			},
		},
	}
}
