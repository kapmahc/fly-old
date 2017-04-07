package site

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"time"

	"github.com/kapmahc/sky"
)

func (p *Engine) renderFuncMap() template.FuncMap {
	return template.FuncMap{
		"uf": p.Layout.URLFor,
		"t":  p.I18n.T,
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"dict": func(values ...interface{}) (sky.H, error) {
			dict := sky.H{}
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict key must be string")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
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
}
