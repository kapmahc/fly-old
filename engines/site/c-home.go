package site

import (
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getLocales(c *gin.Context) (interface{}, error) {
	items, err := p.I18n.Items(c.Param("lang"))
	return items, err
}

func (p *Engine) getSiteInfo(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	data := gin.H{
		"locale":    lang,
		"languages": viper.GetStringSlice("languages"),
	}
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		data[k] = p.I18n.T(lang, "site."+k)
	}

	author := gin.H{}
	for _, k := range []string{"name", "email"} {
		author[k] = p.I18n.T(lang, "site.author."+k)
	}
	data["author"] = author

	return data, nil
}

func (p *Engine) getDashboard(c *gin.Context) (interface{}, error) {
	var items []*web.Dropdown
	web.Walk(func(en web.Engine) error {
		items = append(items, en.Dashboard(c))
		return nil
	})
	return items, nil
}
