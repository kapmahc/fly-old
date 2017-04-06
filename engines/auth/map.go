package auth

import (
	"crypto/aes"
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	_redis "github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/cache/redis"
	s_orm "github.com/kapmahc/sky/i18n/orm"
	"github.com/kapmahc/sky/job"
	"github.com/kapmahc/sky/job/rabbitmq"
	i_orm "github.com/kapmahc/sky/settings/orm"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

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
		&inject.Object{Value: mux.NewRouter()},
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
