package auth

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
	"github.com/unrolled/render"
)

func newRender(rt *mux.Router, i18n *web.I18n, theme string) (*render.Render, error) {
	assets := make(map[string]string)
	if err := filepath.Walk(
		path.Join("themes", theme, "assets"),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			name := info.Name()

			switch filepath.Ext(name) {
			case ".css", ".js", ".png":
				ss := strings.Split(name, ".")
				if len(ss) == 3 {
					assets[fmt.Sprintf("%s.%s", ss[0], ss[2])] = name
				}
			}
			return nil
		},
	); err != nil {
		return nil, err
	}
	for k, v := range assets {
		log.Debugf("assets %-16s => %s", k, v)
	}
	// -------
	funcs := template.FuncMap{
		"t": func(lang, code string, args ...interface{}) string {
			return i18n.T(lang, code, args...)
		},
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"asset": func(k string) string {
			return fmt.Sprintf("/%s", assets[k])
		},
		"str2htm": func(s string) template.HTML {
			return template.HTML(s)
		},
		"fmt": fmt.Sprintf,
		"eq": func(arg1, arg2 interface{}) bool {
			return arg1 == arg2
		},
		"dtf": func(t time.Time) string {
			return t.Format("Mon Jan _2 15:04:05 2006")
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
	}

	// -------
	return render.New(render.Options{
		Directory:  path.Join("themes", theme, "views"),
		IndentJSON: !web.IsProduction(),
		IndentXML:  !web.IsProduction(),
		Layout:     "application",
		Extensions: []string{".html"},
		Funcs:      []template.FuncMap{funcs},
	}), nil

}
