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

	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		return tpl, err
	}
	for i := range items {
		u := &items[i]
		for _, d := range domains {
			if d.ID == u.DomainID {
				u.Domain = d
				break
			}
		}
	}

	data["items"] = items
	return tpl, nil
}

type fmAlias struct {
	Source      string `form:"source" binding:"required,max=255"`
	Destination string `form:"destination" binding:"required,max=255"`
}

func (p *Engine) createAlias(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "ops-mail-aliases-new"

	var users []User
	if err := p.Db.Select([]string{"email", "full_name"}).Order("full_name ASC").Find(&users).Error; err != nil {
		return tpl, err
	}
	data["users"] = users

	if c.Request.Method == http.MethodPost {
		var fm fmAlias
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		var user User
		if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&Alias{
			Source:      fm.Source,
			Destination: fm.Destination,
			DomainID:    user.DomainID,
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

	var users []User
	if err := p.Db.Select([]string{"email", "full_name"}).Order("full_name ASC").Find(&users).Error; err != nil {
		return tpl, err
	}
	data["users"] = users

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
		var user User
		if err := p.Db.Where("email = ?", fm.Destination).First(&user).Error; err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&Alias{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"domain_id":   user.DomainID,
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
