package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexLeaveWords(c *gin.Context) {
	var items []LeaveWord
	err := p.Db.Order("created_at DESC").Find(&items).Error
	web.JSON(c, items, err)
}

type fmLeaveWord struct {
	Body string `form:"body" validate:"required,max=800"`
	Type string `form:"type" validate:"required,max=8"`
}

func (p *Engine) createLeaveWord(c *gin.Context) {
	var fm fmLeaveWord
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Create(&LeaveWord{Type: fm.Type, Body: fm.Body}).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) destroyLeaveWord(c *gin.Context) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error
	web.JSON(c, nil, err)
}
