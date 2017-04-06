package auth

import (
	"crypto/aes"
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	_redis "github.com/garyburd/redigo/redis"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/cache/redis"
	"github.com/kapmahc/sky/i18n"
	s_orm "github.com/kapmahc/sky/i18n/orm"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/job/rabbitmq"
	"github.com/kapmahc/sky/security"
	i_orm "github.com/kapmahc/sky/settings/orm"
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

// Map map object
func (p *Engine) Map(inj *inject.Graph) error {
	db, err := gorm.Open(viper.GetString("database.driver"), sky.DataSource())
	if err != nil {
		return err
	}
	if !sky.IsProduction() {
		db.LogMode(true)
	}

	if err = db.DB().Ping(); err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))

	// -------------
	var tags []language.Tag
	for _, l := range viper.GetStringSlice("languages") {
		lng, er := language.Parse(l)
		if er != nil {
			return er
		}
		tags = append(tags, lng)
	}
	// -----------
	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}
	// -----------

	return inj.Provide(
		&inject.Object{Value: &redis.Store{}},

		&inject.Object{Value: db},
		&inject.Object{Value: s_orm.New(db)},
		&inject.Object{Value: i_orm.New(db)},
		&inject.Object{Value: &_redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (_redis.Conn, error) {
				c, e := _redis.Dial(
					"tcp",
					fmt.Sprintf(
						"%s:%d",
						viper.GetString("redis.host"),
						viper.GetInt("redis.port"),
					),
				)
				if e != nil {
					return nil, e
				}
				if _, e = c.Do("SELECT", viper.GetInt("redis.db")); e != nil {
					c.Close()
					return nil, e
				}
				return c, nil
			},
			TestOnBorrow: func(c _redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}},

		&inject.Object{Value: language.NewMatcher(tags)},

		&inject.Object{Value: job.New()},
		&inject.Object{Value: rabbitmq.New(
			viper.GetString("app.name"),
			viper.GetString("rabbitmq.host"),
			viper.GetInt("rabbitmq.port"),
			viper.GetString("rabbitmq.user"),
			viper.GetString("rabbitmq.password"),
			viper.GetString("rabbitmq.virtual"),
		)},

		&inject.Object{Value: cip},
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
	)
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
		"virtual":  "flv-dev",
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
