package mail

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAliases(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "ops.mail.aliases.index.title")
	tpl := "ops-mail-aliases-index"
	var items []Alias
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return tpl, err
	}
	data["items"] = items
	return tpl, nil
}

type fmAlias struct {
	Source      string `form:"source" binding:"required,max=255"`
	Destination string `form:"destination" binding:"required,max=255"`
	DomainID    uint   `form:"domainId"`
}

func (p *Engine) createAlias(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "ops-mail-aliases-new"
	if c.Request.Method == http.MethodPost {
		var fm fmAlias
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&Alias{
			Source:      fm.Source,
			Destination: fm.Destination,
			DomainID:    fm.DomainID,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/aliases")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateAlias(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "ops-mail-aliases-edit"
	id := c.Param("id")

	var item Alias
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	if c.Request.Method == http.MethodPost {
		var fm fmAlias
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&Alias{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"domain_id":   fm.DomainID,
				"source":      fm.Source,
				"destination": fm.Destination,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/ops/mail/aliases")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyAlias(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Alias{}).Error
	return gin.H{}, err
}
