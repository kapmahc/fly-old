package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexNotices(c *gin.Context) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	web.JSON(c, items, err)
}

type fmNotice struct {
	Body string `form:"body" validate:"required,max=800"`
	Type string `form:"type" validate:"required,max=8"`
}

func (p *Engine) createNotice(c *gin.Context) {
	var fm fmNotice
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Create(&Notice{Type: fm.Type, Body: fm.Body}).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) updateNotice(c *gin.Context) {
	var fm fmNotice
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Model(&Notice{}).Where("id = ?", c.Param("id")).Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) destroyNotice(c *gin.Context) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error
	web.JSON(c, nil, err)
}
