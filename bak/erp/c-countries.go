package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexCountries(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.countries.index.title")
	tpl := "erp-countries-index"
	var items []shop.Country
	if err := p.Db.Select([]string{"id", "name"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmCountry struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createCountry(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-countries-new"
	if c.Request.Method == http.MethodPost {
		var fm fmCountry
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&shop.Country{
			Name: fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/countries")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateCountry(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-countries-edit"
	id := c.Param("id")

	var item shop.Country
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmCountry
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&shop.Country{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"name": fm.Name,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/countries")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyCountry(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(shop.Country{}).Error
	return gin.H{}, err
}
