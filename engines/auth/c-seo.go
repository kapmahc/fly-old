package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
)

func (p *Engine) home(w http.ResponseWriter, r *http.Request) {
	se := viper.GetString("server.home")
	done := false

	web.Walk(func(en web.Engine) error {
		if !done {
			name := strings.ToLower(reflect.ValueOf(en).Elem().Type().String())
			if name == se {
				rt := p.Router.Get(fmt.Sprintf("%s.home", se))
				if rt != nil {
					rt.GetHandler().ServeHTTP(w, r)
					done = true
				}
			}
		}
		return nil
	})
	if !done {
		p.indexNotices(w, r)
	}
}
