package web

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func invoke(m *martini.ClassicMartini, h interface{}) error {
	args, err := m.Invoke(h)
	if err != nil {
		return err
	}
	if len(args) > 0 {
		return args[len(args)-1].Interface().(error)
	}
	return nil
}

// Action ioc action
func Action(fn func(*cli.Context) interface{}) cli.ActionFunc {
	return CfgAction(func(ctx *cli.Context) error {
		mrt := martini.Classic()
		if err := Walk(func(en Engine) error {
			return invoke(mrt, en.Map())
		}); err != nil {
			return err
		}
		return invoke(mrt, fn(ctx))
	})
}

// CfgAction cfg action
func CfgAction(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		log.Infof("read config from config.toml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}
