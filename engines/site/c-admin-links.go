package site

import (
	"net/http"

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

func (p *Engine) indexAdminLinks(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.links.index.title")
	tpl := "site-admin-links-index"
	var links []web.Link
	if err := p.Db.Order("loc ASC, sort_order DESC").Find(&links).Error; err != nil {
		return tpl, err
	}
	data["items"] = links
	return tpl, nil
}

type fmLink struct {
	Label     string `form:"label" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=16"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Engine) createAdminLink(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	data["locs"] = linkLocs
	data["sortOrders"] = sortOrders
	tpl := "site-admin-links-new"
	if c.Request.Method == http.MethodPost {
		var fm fmLink
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&web.Link{
			Loc:       fm.Loc,
			Label:     fm.Label,
			Href:      fm.Href,
			SortOrder: fm.SortOrder,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/links")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateAdminLink(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	data["locs"] = linkLocs
	data["sortOrders"] = sortOrders
	tpl := "site-admin-links-edit"
	id := c.Param("id")

	var item web.Link
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item
	if c.Request.Method == http.MethodPost {
		var fm fmLink
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&web.Link{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"loc":        fm.Loc,
				"href":       fm.Href,
				"sort_order": fm.SortOrder,
				"label":      fm.Label,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/links")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyAdminLink(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(web.Link{}).Error
	return gin.H{}, err
}
