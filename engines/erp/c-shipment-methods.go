package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexShippingMethods(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.shipping-methods.index.title")
	tpl := "erp-shipping-methods-index"
	var items []shop.ShippingMethod
	if err := p.Db.Select([]string{"id", "name", "logo", "active"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmShippingMethod struct {
	Name        string `form:"name" binding:"required,max=255"`
	Logo        string `form:"logo" binding:"required,max=255"`
	Tracking    string `form:"tracking" binding:"required,max=255"`
	Description string `form:"description" binding:"required"`
	Active      bool   `form:"active"`
}

func (p *Engine) createShippingMethod(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-shipping-methods-new"
	if c.Request.Method == http.MethodPost {
		var fm fmShippingMethod
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&shop.ShippingMethod{
			Logo:        fm.Logo,
			Tracking:    fm.Tracking,
			Description: fm.Description,
			Active:      fm.Active,
			Name:        fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/shipping-methods")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateShippingMethod(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-shipping-methods-edit"
	id := c.Param("id")

	var item shop.ShippingMethod
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmShippingMethod
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&shop.ShippingMethod{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"description": fm.Description,
				"active":      fm.Active,
				"name":        fm.Name,
				"tracking":    fm.Tracking,
				"logo":        fm.Logo,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/shipping-methods")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyShippingMethod(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(shop.ShippingMethod{}).Error
	return gin.H{}, err
}
