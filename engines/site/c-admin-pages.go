package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

var (
	pageLocs = []interface{}{"carousel", "circle", "square"}
	defLogo  = "data:image/gif;base64,R0lGODlhAQABAIAAAHd3dwAAACH5BAAAAAAALAAAAAABAAEAAAICRAEAOw=="
)

func (p *Engine) indexAdminPages(c *gin.Context) (interface{}, error) {
	var pages []web.Page
	err := p.Db.Order("loc ASC, sort_order DESC").Find(&pages).Error
	return pages, err
}

type fmPage struct {
	Title     string `form:"title" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	Action    string `form:"action" binding:"required,max=32"`
	Logo      string `form:"logo" binding:"required,max=255"`
	Summary   string `form:"summary" binding:"required,max=2048"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Engine) createAdminPage(c *gin.Context) (interface{}, error) {

	var fm fmPage
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Create(&web.Page{
		Loc:       fm.Loc,
		Title:     fm.Title,
		Summary:   fm.Summary,
		Logo:      fm.Logo,
		Href:      fm.Href,
		Action:    fm.Action,
		SortOrder: fm.SortOrder,
	}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) updateAdminPage(c *gin.Context) (interface{}, error) {
	id := c.Param("id")

	var item web.Page
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	var fm fmPage
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&web.Page{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
			"title":      fm.Title,
			"summary":    fm.Summary,
			"action":     fm.Action,
			"logo":       fm.Logo,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Engine) destroyAdminPage(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Page{}).Error
	return gin.H{}, err
}
