package auth

import (
	"context"
	"net/http"
	"reflect"
	"strings"

	log "github.com/Sirupsen/logrus"
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
	log.Debug(data)
	next(w, r.WithContext(ctx))
}
