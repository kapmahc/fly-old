package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminLocales(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var items []web.Locale
	err := p.Db.Select([]string{"id", "code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").Find(&items).Error
	return items, err
}

func (p *Engine) deleteAdminLocales(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(&web.Locale{}).Error

	return gin.H{}, err
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

func (p *Engine) postAdminLocales(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	var fm fmLocale
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.I18n.Set(lang, fm.Code, fm.Message); err != nil {
		return nil, err
	}

	var l web.Locale
	err := p.Db.Where("lang = ? AND code = ?", lang, fm.Code).First(&l).Error
	return l, err
}
