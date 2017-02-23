package auth

import (
	"net/http"
	"reflect"

	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
)

func (p *Engine) home(w http.ResponseWriter, r *http.Request) {
	se := viper.GetString("server.home")
	web.Walk(func(en web.Engine) error {
		if se == reflect.ValueOf(en).Elem().Type().PkgPath() {
			en.Home()(w, r)
		}
		return nil
	})
}
