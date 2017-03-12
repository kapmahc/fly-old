package forum

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexTags(c *gin.Context) {
	var tags []Tag
	err := p.Db.Find(&tags).Error
	web.JSON(c, tags, err)
}

type fmTag struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createTag(c *gin.Context) {

	var fm fmTag
	err := c.Bind(&fm)
	var t *Tag
	if err == nil {
		t = &Tag{Name: fm.Name}
		err = p.Db.Create(t).Error
	}
	web.JSON(c, t, err)
}

func (p *Engine) showTag(c *gin.Context) {
	var tag Tag
	err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error
	if err == nil {
		err = p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error
	}
	web.JSON(c, tag, err)
}

func (p *Engine) updateTag(c *gin.Context) {
	var fm fmTag
	var tag Tag
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Where("id = ?", c.Param("id")).First(&tag).Error
	}
	if err == nil {
		err = p.Db.Model(&tag).Update("name", fm.Name).Error
	}
	web.JSON(c, tag, err)
}

func (p *Engine) destroyTag(c *gin.Context) {
	var tag Tag
	err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error
	if err == nil {
		err = p.Db.Model(&tag).Association("Articles").Clear().Error
	}
	if err == nil {
		err = p.Db.Delete(&tag).Error
	}
	web.JSON(c, nil, err)
}
