package auth

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/i18n"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/security"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Engine engine
type Engine struct {
	Dao    *Dao             `inject:""`
	I18n   *i18n.I18n       `inject:""`
	Db     *gorm.DB         `inject:""`
	Cipher *security.Cipher `inject:""`
}

// Mount web mount points
func (p *Engine) Mount(*sky.Router) {

}

// Workers job workers
func (p *Engine) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Atom rss.atom
func (p *Engine) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Engine) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

// Navbar navbar
func (p *Engine) Navbar(*sky.Context) []*sky.Dropdown {
	return []*sky.Dropdown{}
}

// Dashboard dashboard
func (p *Engine) Dashboard(*sky.Context) []*sky.Dropdown {
	return []*sky.Dropdown{}
}

func init() {
	viper.SetEnvPrefix("sky")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetDefault("app", map[string]interface{}{
		"name": "fly",
	})
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
		"port":   3000,
		"ssl":    true,
		"themes": []string{"bootstrap"},
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
