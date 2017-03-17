package web

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/text/language"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var (
	// Version version
	Version string
	// BuildTime build time
	BuildTime string
)

// Main main entry
func Main() error {

	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Version = fmt.Sprintf("%s(%s)", Version, BuildTime)
	app.Usage = "FLY - A complete open source e-commerce solution by the Go language and React."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{}

	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	return app.Run(os.Args)
}

func init() {
	viper.SetEnvPrefix("fly")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":    "localhost",
			"port":    5432,
			"user":    "postgres",
			"dbname":  "fly_dev",
			"sslmode": "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"name":     "fly",
		"port":     8080,
		"backend":  "http://localhost:8080",
		"frontend": "http://localhost:3000",
		"ssl":      true,
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":  RandomStr(32),
		"aes":  RandomStr(32),
		"hmac": RandomStr(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	viper.SetDefault("languages", []string{
		language.AmericanEnglish.String(),
		language.SimplifiedChinese.String(),
		language.TraditionalChinese.String(),
	})
}
