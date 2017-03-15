package forum

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) indexTags(c *gin.Context) (interface{}, error) {
	var tags []Tag
	err := p.Db.Find(&tags).Error
	return tags, err
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
	err := p.Db.Create(&t).Error
	return t, err
}

func (p *Engine) showTag(c *gin.Context) (interface{}, error) {
	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return nil, err
	}

	err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error
	return tag, err
}

func (p *Engine) updateTag(c *gin.Context) (interface{}, error) {
	var fm fmTag
	var tag Tag
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return nil, err
	}
	err := p.Db.Model(&tag).Update("name", fm.Name).Error
	return tag, err
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
