package site

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) openRender(theme string) (*template.Template, error) {
	// assets := make(map[string]string)
	// if err := filepath.Walk(
	// 	path.Join("themes", theme, "public", "assets"),
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if info.IsDir() {
	// 			return nil
	// 		}
	// 		name := info.Name()
	//
	// 		switch filepath.Ext(name) {
	// 		case ".css", ".js", ".png", ".svg":
	// 			ss := strings.Split(name, ".")
	// 			if len(ss) == 3 {
	// 				assets[fmt.Sprintf("%s.%s", ss[0], ss[2])] = name
	// 			}
	// 		}
	// 		return nil
	// 	},
	// ); err != nil {
	// 	return nil, err
	// }
	// for k, v := range assets {
	// 	log.Debugf("assets %-16s => %s", k, v)
	// }

	// ----------------

	funcs := template.FuncMap{
		"t": p.I18n.T,
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"dict": func(values ...interface{}) (gin.H, error) {
			dict := gin.H{}
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		// "asset": func(k string) string {
		// 	if web.IsProduction() {
		// 		return fmt.Sprintf("/assets/%s", assets[k])
		// 	}
		// 	return fmt.Sprintf("/public/assets/%s", assets[k])
		// },
		"even": func(i interface{}) bool {
			if i != nil {
				switch i.(type) {
				case int:
					return i.(int)%2 == 0
				case uint:
					return i.(uint)%2 == 0
				case int64:
					return i.(int64)%2 == 0
				case uint64:
					return i.(uint64)%2 == 0
				}
			}
			return false
		},
		"fmt": fmt.Sprintf,
		"eq": func(arg1, arg2 interface{}) bool {
			return arg1 == arg2
		},
		"str2htm": func(s string) template.HTML {
			return template.HTML(s)
		},
		"dtf": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 15:04:05 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"df": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"links": func(loc string) []web.Link {
			var items []web.Link
			if err := p.Db.Where("loc = ?", loc).Order("sort_order DESC").Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"pages": func(loc string) []web.Page {
			var items []web.Page
			if err := p.Db.Where("loc = ?", loc).Order("sort_order DESC").Find(&items).Error; err != nil {
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
		"starts": func(s string, b string) bool {
			return strings.HasPrefix(s, b)
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
