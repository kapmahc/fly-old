package site

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexLeaveWords(c *gin.Context) (interface{}, error) {
	var items []LeaveWord
	err := p.Db.Order("created_at DESC").Find(&items).Error
	return items, err
}

type fmLeaveWord struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) createLeaveWord(c *gin.Context) (interface{}, error) {
	var fm fmLeaveWord
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	err := p.Db.Create(&LeaveWord{Type: fm.Type, Body: fm.Body}).Error
	return gin.H{}, err
}

func (p *Engine) destroyLeaveWord(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error
	return gin.H{}, err
}
