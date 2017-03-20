package site

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAdminNotices(c *gin.Context, lang string, data gin.H) (string, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	data["items"] = items
	data["title"] = p.I18n.T(lang, "site.notices.index.title")
	return "site-notices-manage", err
}

func (p *Engine) indexNotices(c *gin.Context, lang string, data gin.H) (string, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	data["items"] = items
	data["title"] = p.I18n.T(lang, "site.notices.index.title")
	return "site-notices-index", err
}

type fmNotice struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) createNotice(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "site-notices-form"
	if c.Request.Method == http.MethodPost {
		var fm fmNotice
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		n := Notice{Type: fm.Type, Body: fm.Body}
		if err := p.Db.Create(&n).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/notices")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateNotice(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "site-notices-form"
	id := c.Param("id")

	var n Notice
	if err := p.Db.Where("id = ?", id).First(&n).Error; err != nil {
		return tpl, err
	}
	data["body"] = n.Body

	if c.Request.Method == http.MethodPost {
		var fm fmNotice
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&Notice{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"body": fm.Body,
				"type": fm.Type,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/notices")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyNotice(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error
	return gin.H{}, err
}
