package site

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

var (
	linkLocs   = []interface{}{"top"}
	sortOrders = []interface{}{}
)

func init() {
	for i := -10; i <= 10; i++ {
		sortOrders = append(sortOrders, i)
	}
}

func (p *Engine) indexAdminLinks(c *gin.Context) (interface{}, error) {
	var links []web.Link
	err := p.Db.Order("loc ASC, sort_order DESC").Find(&links).Error
	return links, err
}

type fmLink struct {
	Label     string `form:"label" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Engine) createAdminLink(c *gin.Context) (interface{}, error) {

	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	item := web.Link{
		Loc:       fm.Loc,
		Label:     fm.Label,
		Href:      fm.Href,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (p *Engine) updateAdminLink(c *gin.Context) (interface{}, error) {
	id := c.Param("id")

	var item web.Link
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	var fm fmLink
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(&web.Link{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
			"label":      fm.Label,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyAdminLink(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Link{}).Error
	return gin.H{}, err
}
