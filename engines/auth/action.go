package auth

import (
	"crypto/aes"

	"golang.org/x/text/language"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gorilla/sessions"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type injectLogger struct {
}

func (p *injectLogger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Action ioc action
func Action(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return web.Action(func(ctx *cli.Context) error {
		inj := inject.Graph{Logger: &injectLogger{}}
		// -------
		var tags []language.Tag
		for _, l := range viper.GetStringSlice("languages") {
			if lng, err := language.Parse(l); err == nil {
				tags = append(tags, lng)
			} else {
				return err
			}
		}
		// -------------------
		db, err := web.OpenDatabase()
		if err != nil {
			return err
		}
		// -------------------
		rep := web.OpenRedis()
		// -------------------
		bws, err := web.NewWorkerServer()
		if err != nil {
			return err
		}
		// -------------------
		cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
		if err != nil {
			return err
		}
		// -----------
		var i18n web.I18n
		rdr, err := openRender(viper.GetString("server.theme"), &i18n)
		if err != nil {
			return err
		}

		if err := inj.Provide(
			&inject.Object{Value: db},
			&inject.Object{Value: bws},
			&inject.Object{Value: rep},
			&inject.Object{Value: rdr},
			&inject.Object{Value: &i18n},
			&inject.Object{Value: language.NewMatcher(tags)},
			&inject.Object{Value: cip, Name: "aes.cip"},
			&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
			&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
			&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
			&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
			&inject.Object{Value: sessions.NewCookieStore([]byte("secrets.session"))},
		); err != nil {
			return err
		}
		// -----------------
		if err := web.Walk(func(en web.Engine) error {
			return inj.Provide(&inject.Object{Value: en})
		}); err != nil {
			return err
		}

		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx, &inj)
	})
}
