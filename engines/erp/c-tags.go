package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/mall"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexTags(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.tags.index.title")
	tpl := "erp-tags-index"
	var tags []mall.Tag
	if err := p.Db.Order("updated_at DESC").Find(&tags).Error; err != nil {
		return tpl, err
	}
	data["items"] = tags
	return tpl, nil
}

type fmTag struct {
	Name        string `form:"name" binding:"required,max=255"`
	Description string `form:"description" binding:"required"`
}

func (p *Engine) createTag(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-tags-new"
	if c.Request.Method == http.MethodPost {
		var fm fmTag
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&mall.Tag{
			Model: mall.Model{
				Name:        fm.Name,
				Description: fm.Description,
			},
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/tags")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateTag(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-tags-edit"
	id := c.Param("id")

	var tag mall.Tag
	if err := p.Db.Where("id = ?", id).First(&tag).Error; err != nil {
		return tpl, err
	}
	data["item"] = tag

	switch c.Request.Method {
	case http.MethodPost:
		var fm fmTag
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&mall.Tag{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"name":        fm.Name,
				"description": fm.Description,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/tags")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyTag(c *gin.Context) (interface{}, error) {
	err := p.Db.Where("id = ?", c.Param("id")).Delete(mall.Tag{}).Error
	return gin.H{}, err
}
