package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexZones(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.zones.index.title")
	tpl := "erp-zones-index"
	var items []shop.Zone
	if err := p.Db.Select([]string{"id", "name", "active"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmZone struct {
	Name   string `form:"name" binding:"required,max=255"`
	Active bool   `form:"active"`
}

func (p *Engine) createZone(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-zones-new"
	if c.Request.Method == http.MethodPost {
		var fm fmZone
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&shop.Zone{
			Active: fm.Active,
			Name:   fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/zones")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateZone(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-zones-edit"
	id := c.Param("id")

	var item shop.Zone
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmZone
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&shop.Zone{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"active": fm.Active,
				"name":   fm.Name,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/zones")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyZone(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(shop.Zone{}).Error
	return gin.H{}, err
}
