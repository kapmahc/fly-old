package auth

import (
	"crypto/aes"
	"fmt"
	"html/template"
	"path"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
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
		// -----------------
		i18n, err := web.NewI18n()
		if err != nil {
			return err
		}
		// -------------------
		rt := mux.NewRouter()
		// -------------------
		rdr := render.New(render.Options{
			Directory:  path.Join("themes", viper.GetString("server.theme"), "views"),
			IndentJSON: !web.IsProduction(),
			IndentXML:  !web.IsProduction(),
			Layout:     "application",
			Extensions: []string{".html"},
			Funcs: []template.FuncMap{
				{
					"t": func(lang, code string, args ...interface{}) string {
						return i18n.T(lang, code, args...)
					},
					"f": func(code string, args ...interface{}) string {
						return fmt.Sprintf(code, args...)
					},
					"assets_css": func(name string) template.HTML {
						return template.HTML(fmt.Sprintf(`<link rel="stylesheet" href="/css/%s.css">`, name))
					},
					"assets_js": func(name string) template.HTML {
						return template.HTML(fmt.Sprintf(`<script src="/js/%s.js"></script>`, name))
					},
					"str2htm": func(s string) template.HTML {
						return template.HTML(s)
					},
					"uf": func(name string, args ...interface{}) string {
						var pairs []string
						for _, arg := range args {
							pairs = append(pairs, fmt.Sprintf("%v", arg))
						}
						url, err := rt.Get(name).URL(pairs...)
						if err != nil {
							return err.Error()
						}
						return url.String()
					},
				},
			},
		})

		if err := inj.Provide(
			&inject.Object{Value: db},
			&inject.Object{Value: bws},
			&inject.Object{Value: rep},
			&inject.Object{Value: cip},
			&inject.Object{Value: i18n},
			&inject.Object{Value: rdr},
			&inject.Object{Value: rt},
			&inject.Object{Value: cip, Name: "aes.cip"},
			&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
			&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
			&inject.Object{Value: viper.GetString("app.name"), Name: "namespace"},
			&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},
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
