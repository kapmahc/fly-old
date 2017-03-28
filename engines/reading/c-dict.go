package reading

import (
	"html/template"
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmDict struct {
	Keywords string `form:"keywords" binding:"required,max=255"`
}

func (p *Engine) formDict(c *gin.Context) (interface{}, error) {
	data["title"] = p.I18n.T(lang, "reading.dict.index.title")
	tpl := "reading-dict-index"
	switch c.Request.Method {
	case http.MethodPost:
		var fm fmDict
		if err := c.Bind(&fm); err != nil {
			return nil, err
		}
		data["keywords"] = fm.Keywords
		rst := gin.H{}
		for _, dic := range dictionaries {
			for _, sen := range dic.Translate(fm.Keywords) {
				var items []interface{}
				for _, pat := range sen.Parts {
					switch pat.Type {
					case 'g', 'h':
						items = append(items, template.HTML(pat.Data))
					default:
						items = append(items, string(pat.Data))
					}
				}
				rst[dic.GetBookName()] = items
			}
		}
		data["result"] = rst
	}
	return tpl, nil
}
