package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexPaymentMethods(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.payment-methods.index.title")
	tpl := "erp-payment-methods-index"
	var items []shop.PaymentMethod
	if err := p.Db.Select([]string{"id", "name", "type", "active"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmPaymentMethod struct {
	Name        string `form:"name" binding:"required,max=255"`
	Type        string `form:"type" binding:"required,max=16"`
	Description string `form:"description" binding:"required"`
	Active      bool   `form:"active"`
}

func (p *Engine) createPaymentMethod(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-payment-methods-new"
	if c.Request.Method == http.MethodPost {
		var fm fmPaymentMethod
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&shop.PaymentMethod{
			Type:        fm.Type,
			Description: fm.Description,
			Active:      fm.Active,
			Name:        fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/payment-methods")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updatePaymentMethod(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-payment-methods-edit"
	id := c.Param("id")

	var item shop.PaymentMethod
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmPaymentMethod
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&shop.PaymentMethod{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"description": fm.Description,
				"active":      fm.Active,
				"name":        fm.Name,
				"type":        fm.Type,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/payment-methods")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyPaymentMethod(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(shop.PaymentMethod{}).Error
	return gin.H{}, err
}
