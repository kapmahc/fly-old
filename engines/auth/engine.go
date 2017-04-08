package auth

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/i18n"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/security"
	"github.com/kapmahc/sky/settings"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Dao      *Dao               `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Db       *gorm.DB           `inject:""`
	Cipher   *security.Cipher   `inject:""`
	Layout   *Layout            `inject:""`
	Queue    job.Queue          `inject:""`
	Jwt      *Jwt               `inject:""`
	Server   *job.Server        `inject:""`
	Settings *settings.Settings `inject:""`
	Hmac     *security.Hmac     `inject:""`
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Application application
func (p *Engine) Application(c *sky.Context) []*sky.Dropdown {
	lang := c.Get(sky.LOCALE).(string)
	return []*sky.Dropdown{
		&sky.Dropdown{Label: p.I18n.T(lang, "auth.users.index.title"), Href: p.Layout.URLFor("auth.users.index")},
	}
}

// Dashboard dashboard
func (p *Engine) Dashboard(c *sky.Context) []*sky.Dropdown {
	lang := c.Get(sky.LOCALE).(string)

	var items []*sky.Dropdown
	if _, ok := c.Get(CurrentUser).(*User); ok {
		items = append(
			items,
			&sky.Dropdown{
				Label: p.I18n.T(lang, "auth.dashboard.title"),
				Links: []*sky.Link{
					&sky.Link{Label: p.I18n.T(lang, "auth.users.logs.title"), Href: p.Layout.URLFor("auth.users.logs")},
					nil,
					&sky.Link{Label: p.I18n.T(lang, "auth.users.info.title"), Href: p.Layout.URLFor("auth.users.info")},
					&sky.Link{Label: p.I18n.T(lang, "auth.users.change-password.title"), Href: p.Layout.URLFor("auth.users.change-password")},
				},
			},
		)
	}
	return items
}

func init() {
	viper.SetEnvPrefix("sky")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})
	viper.SetDefault("rabbitmq", map[string]interface{}{
		"user":     "guest",
		"password": "guest",
		"host":     "localhost",
		"port":     "5672",
		"virtual":  "fly-dev",
	})
	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"user":     "postgres",
			"password": "",
			"dbname":   "fly_dev",
			"sslmode":  "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"name":  "www.change-me.com",
		"port":  3000,
		"ssl":   true,
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":    sky.Random(32),
		"aes":    sky.Random(32),
		"hmac":   sky.Random(32),
		"csrf":   sky.Random(32),
		"cookie": sky.Random(32),
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

	// ------------
	sky.Register(&Engine{})
}
