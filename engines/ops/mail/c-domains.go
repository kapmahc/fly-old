package mail

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexDomains(c *gin.Context) (interface{}, error) {

	var items []Domain
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

type fmDomain struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createDomain(c *gin.Context) (interface{}, error) {

	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	item := Domain{
		Name: fm.Name,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil

}

func (p *Engine) updateDomain(c *gin.Context) (interface{}, error) {

	id := c.Param("id")
	var fm fmDomain
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&Domain{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name": fm.Name,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyDomain(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Domain{}).Error
	return gin.H{}, err
}
