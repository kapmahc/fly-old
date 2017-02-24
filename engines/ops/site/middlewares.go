package site

import (
	"context"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) layoutMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data := r.Context().Value(web.DATA).(web.H)
	ctx := context.WithValue(r.Context(), web.DATA, data)
	var engines []string
	web.Walk(func(en web.Engine) error {
		engines = append(engines, strings.ToLower(reflect.ValueOf(en).Elem().Type().String()))
		return nil
	})
	data["engines"] = engines
	data[csrf.TemplateTag] = csrf.TemplateField(r)
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	next(w, r.WithContext(ctx))
}
