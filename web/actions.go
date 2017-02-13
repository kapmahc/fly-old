package web

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// Action cfg action
func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		log.Infof("read config from config.toml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
