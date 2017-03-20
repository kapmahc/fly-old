package erp

import (
	"net/http"

	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexStates(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.states.index.title")
	tpl := "erp-states-index"
	states, err := p.Dao.ListState()
	if err != nil {
		return tpl, err
	}
	zones, err := p.Dao.ListZone()
	if err != nil {
		return tpl, err
	}
	countries, err := p.Dao.ListCountry()
	if err != nil {
		return tpl, err
	}
	for i := range states {
		s := &states[i]
		for _, c := range countries {
			if c.ID == s.CountryID {
				s.Country = c
			}
		}
		for _, z := range zones {
			if z.ID == s.ZoneID {
				s.Zone = z
			}
		}
	}
	data["states"] = states
	return tpl, nil
}

type fmState struct {
	Name      string `form:"name" binding:"required,max=255"`
	ZoneID    uint   `form:"zoneId"`
	CountryID uint   `form:"countryId"`
}

func (p *Engine) createState(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "erp-states-new"
	var err error
	if data["countries"], err = p.Dao.ListCountry(); err != nil {
		return tpl, err
	}
	if data["zones"], err = p.Dao.ListZone(); err != nil {
		return tpl, err
	}
	if c.Request.Method == http.MethodPost {
		var fm fmState
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Create(&shop.State{
			CountryID: fm.CountryID,
			ZoneID:    fm.ZoneID,
			Name:      fm.Name,
		}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/states")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) updateState(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "erp-states-edit"
	id := c.Param("id")

	var item shop.State
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return tpl, err
	}
	data["item"] = item

	var err error
	if data["countries"], err = p.Dao.ListCountry(); err != nil {
		return tpl, err
	}
	if data["zones"], err = p.Dao.ListZone(); err != nil {
		return tpl, err
	}

	if c.Request.Method == http.MethodPost {
		var fm fmState
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		if err := p.Db.Model(&shop.State{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"name":       fm.Name,
				"zone_id":    fm.ZoneID,
				"country_id": fm.CountryID,
			}).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/erp/states")
		return "", nil
	}

	return tpl, nil
}

func (p *Engine) destroyState(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(shop.State{}).Error
	return gin.H{}, err
}
