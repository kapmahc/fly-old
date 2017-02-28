package site

import (
	"context"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) layoutMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data := r.Context().Value(web.DATA).(web.H)
	ctx := context.WithValue(r.Context(), web.DATA, data)
	// var engines []string
	// web.Walk(func(en web.Engine) error {
	// 	engines = append(engines, strings.ToLower(reflect.ValueOf(en).Elem().Type().String()))
	// 	return nil
	// })
	// data["engines"] = engines

	var navbar []web.Dropdown
	var dashboard []web.Link
	web.Walk(func(en web.Engine) error {
		nb, db := en.NavBar(r)
		navbar = append(navbar, nb...)
		dashboard = append(dashboard, db...)
		return nil
	})
	data["navbar"] = navbar
	data["dashboard"] = dashboard

	data[csrf.TemplateTag] = csrf.TemplateField(r)
	tkn := csrf.Token(r)
	data["csrf"] = tkn
	w.Header().Set("X-CSRF-Token", tkn)
	next(w, r.WithContext(ctx))
}
