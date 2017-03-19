package site

import (
	"net/http"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAdminPages(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.pages.index.title")
	tpl := "site-admin-pages-index"
	var pages []web.Page
	if err := p.Db.Order("loc ASC, sort_order DESC").Find(&pages).Error; err != nil {
		return tpl, err
	}
	data["items"] = pages
	return tpl, nil
}

type fmPage struct {
	Title     string `form:"title" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	Logo      string `form:"logo" binding:"required,max=255"`
	Summary   string `form:"summary" binding:"required,max=2048"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Engine) createAdminPage(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "site-admin-pages-new"
	if c.Request.Method == http.MethodPost {
		var fm fmPage
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&web.Page{
			Loc:       fm.Loc,
			Title:     fm.Title,
			Summary:   fm.Summary,
			Logo:      fm.Logo,
			Href:      fm.Href,
			SortOrder: fm.SortOrder,
		}).Error; err != nil {
			return tpl, err
		}
		c.JSON(http.StatusFound, "/admin/pages")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateAdminPage(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "site-admin-pages-edit"
	id := c.Param("id")

	var item web.Page
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item
	if c.Request.Method == http.MethodPost {
		var fm fmPage
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&web.Page{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"loc":        fm.Loc,
				"href":       fm.Href,
				"sort_order": fm.SortOrder,
				"title":      fm.Title,
				"summary":    fm.Summary,
				"logo":       fm.Logo,
			}).Error; err != nil {
			return tpl, err
		}
		c.JSON(http.StatusFound, "/admin/pages")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyAdminPage(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Page{}).Error
	return gin.H{}, err
}
