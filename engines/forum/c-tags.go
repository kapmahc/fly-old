package forum

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexTags(c *gin.Context) (interface{}, error) {

	var tags []Tag
	if err := p.Db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

type fmTag struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createTag(c *gin.Context) (interface{}, error) {

	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	t := Tag{Name: fm.Name}
	if err := p.Db.Create(&t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (p *Engine) showTag(c *gin.Context) (interface{}, error) {

	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (p *Engine) updateTag(c *gin.Context) (interface{}, error) {

	id := c.Param("id")

	var tag Tag
	if err := p.Db.Where("id = ?", id).First(&tag).Error; err != nil {
		return nil, err
	}
	var fm fmTag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&Tag{}).Where("id = ?", id).Update("name", fm.Name).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyTag(c *gin.Context) (interface{}, error) {
	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&tag).Association("Articles").Clear().Error; err != nil {
		return nil, err
	}

	err := p.Db.Delete(&tag).Error
	return gin.H{}, err
}
