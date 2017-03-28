package site

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexNotices(c *gin.Context) (interface{}, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error

	return items, err
}

type fmNotice struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) createNotice(c *gin.Context) (interface{}, error) {

	var fm fmNotice
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	n := Notice{Type: fm.Type, Body: fm.Body}
	if err := p.Db.Create(&n).Error; err != nil {
		return nil, err
	}

	return n, nil
}

func (p *Engine) updateNotice(c *gin.Context) (interface{}, error) {
	id := c.Param("id")
	var n Notice
	if err := p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return nil, err
	}

	var fm fmNotice
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&Notice{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) destroyNotice(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error
	return gin.H{}, err
}
