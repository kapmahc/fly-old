package auth

import (
	"context"
	"net/http"
	"reflect"

	"github.com/kapmahc/fly/web"
)

func (p *Engine) layoutMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data := r.Context().Value(web.DATA).(web.H)
	ctx := context.WithValue(r.Context(), web.DATA, data)
	var engines []string
	web.Walk(func(en web.Engine) error {
		reflect.ValueOf(en).Elem().Type()
		engines = append(engines, reflect.ValueOf(en).Elem().Type().PkgPath())
		return nil
	})
	data["engines"] = engines
	next(w, r.WithContext(ctx))
}
