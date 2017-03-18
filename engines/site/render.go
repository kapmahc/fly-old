package site

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
	"github.com/kapmahc/fly/web"
)

func (p *Engine) openRender(theme string) (*template.Template, error) {
	assets := make(map[string]string)
	if err := filepath.Walk(
		path.Join("themes", theme, "public", "assets"),
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

	// ----------------

	funcs := template.FuncMap{
		"t": p.I18n.T,
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"asset": func(k string) string {
			return fmt.Sprintf("/public/assets/%s", assets[k])
		},
		"fmt": fmt.Sprintf,
		"eq": func(arg1, arg2 interface{}) bool {
			return arg1 == arg2
		},
		"str2htm": func(s string) template.HTML {
			return template.HTML(s)
		},
		"dtf": func(t time.Time) string {
			return t.Format("Mon Jan _2 15:04:05 2006")
		},
		"links": func(loc string) []web.Link {
			var items []web.Link
			if err := p.Db.Order("sort_order DESC").Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"pages": func(loc string) []web.Page {
			var items []web.Page
			if err := p.Db.Order("sort_order DESC").Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"in": func(o interface{}, args []interface{}) bool {
			for _, v := range args {
				if o == v {
					return true
				}
			}
			return false
		},
	}

	var files []string
	filepath.Walk(path.Join("themes", theme, "views"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})

	return template.New("").Funcs(funcs).ParseFiles(files...)
}
