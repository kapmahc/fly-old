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
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) createNotice(c *gin.Context) {
	var fm fmNotice
	err := c.Bind(&fm)
	var n *Notice
	if err == nil {
		n = &Notice{Type: fm.Type, Body: fm.Body}
		err = p.Db.Create(n).Error
	}
	web.JSON(c, n, err)
}

func (p *Engine) updateNotice(c *gin.Context) {
	var fm fmNotice
	err := c.Bind(&fm)
	var n Notice
	if err == nil {
		err = p.Db.Where("id = ?", c.Param("id")).First(&n).Error
	}
	if err == nil {
		err = p.Db.Model(&n).Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error
	}
	web.JSON(c, n, err)
}

func (p *Engine) destroyNotice(c *gin.Context) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error
	web.JSON(c, nil, err)
}
