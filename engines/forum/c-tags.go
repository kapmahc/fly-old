package forum

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexTags(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "forum.tags.index.title")
	tpl := "forum-tags-index"
	var tags []Tag
	if err := p.Db.Find(&tags).Error; err != nil {
		return tpl, err
	}
	data["items"] = tags
	if c.Request.URL.Query().Get("act") == "manage" {
		if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
			return "forum-tags-manage", nil
		}
		c.AbortWithStatus(http.StatusForbidden)
		return "", nil
	}
	return tpl, nil
}

type fmTag struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createTag(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "forum-tags-new"
	if c.Request.Method == http.MethodPost {
		var fm fmTag
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		t := Tag{Name: fm.Name}
		if err := p.Db.Create(&t).Error; err != nil {
			return tpl, err
		}
	}
	return tpl, nil
}

func (p *Engine) showTag(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "forum.tags.show.title")
	tpl := "forum-tags-show"
	var tag Tag
	if err := p.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		return tpl, err
	}

	if err := p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error; err != nil {
		return tpl, err
	}
	return tpl, nil
}

func (p *Engine) updateTag(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "forum-tags-edit"
	id := c.Param("id")
	switch c.Request.Method {
	case http.MethodGet:
		var tag Tag
		if err := p.Db.Where("id = ?", id).First(&tag).Error; err != nil {
			return tpl, err
		}
		data["name"] = tag.Name
	case http.MethodPost:
		var fm fmTag
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if err := p.Db.Model(&Tag{}).Where("id = ?", id).Update("name", fm.Name).Error; err != nil {
			return tpl, err
		}
		data["name"] = fm.Name
	}
	return tpl, nil
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
