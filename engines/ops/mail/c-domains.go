package mail

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexDomains(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.domains.index.title")
	tpl := "ops-mail-domains-index"
	var items []Domain
	if err := p.Db.Select([]string{"id", "name", "active"}).Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmDomain struct {
	Name string `form:"name" binding:"required,max=255"`
}

func (p *Engine) createDomain(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "ops-mail-domains-new"
	if c.Request.Method == http.MethodPost {
		var fm fmDomain
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&Domain{
			Name: fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/domains")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateDomain(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "ops-mail-domains-edit"
	id := c.Param("id")

	var item Domain
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmDomain
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&Domain{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"name": fm.Name,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/domains")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyDomain(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Domain{}).Error
	return gin.H{}, err
}
