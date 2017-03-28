package mail

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexAliases(c *gin.Context) (interface{}, error) {

	var items []Alias
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		return nil, err
	}
	for i := range items {
		u := &items[i]
		for _, d := range domains {
			if d.ID == u.DomainID {
				u.Domain = d
				break
			}
		}
	}

	return items, nil
}

type fmAlias struct {
	Source      string `form:"source" binding:"required,max=255"`
	Destination string `form:"destination" binding:"required,max=255"`
}

func (p *Engine) createAlias(c *gin.Context) (interface{}, error) {

	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		return nil, err
	}
	item := Alias{
		Source:      fm.Source,
		Destination: fm.Destination,
		DomainID:    user.DomainID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (p *Engine) updateAlias(c *gin.Context) (interface{}, error) {

	id := c.Param("id")

	var fm fmAlias
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&Alias{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"domain_id":   user.DomainID,
			"source":      fm.Source,
			"destination": fm.Destination,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyAlias(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Alias{}).Error
	return gin.H{}, err
}
