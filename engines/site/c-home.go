package site

import (
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getLocales(c *gin.Context) {
	lang, err := language.Parse(c.Param("lang"))
	data := p.I18n.Items(lang.String())
	web.JSON(c, data, err)
}

func (p *Engine) getSiteInfo(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)
	data := gin.H{"locale": lang}
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		data[k] = p.I18n.T(lang, "site."+k)
	}
	author := gin.H{}
	for _, k := range []string{"name", "email"} {
		author[k] = p.I18n.T(lang, "site.author."+k)
	}
	data["author"] = author
	data["languages"] = viper.GetStringSlice("languages")
	web.JSON(c, data, nil)
}
