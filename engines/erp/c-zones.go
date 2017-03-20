package erp

import (
	"github.com/kapmahc/fly/engines/shop"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexZones(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "erp.zones.index.title")
	tpl := "erp-zones-index"
	var states []shop.State
	if err := p.Db.Select([]string{"id", "name", "country_id", "zone_id", "active"}).Order("name Asc").Find(&states).Error; err != nil {
		return tpl, err
	}
	var zones []shop.Zone
	if err := p.Db.Select([]string{"id", "name"}).Order("name Asc").Find(&zones).Error; err != nil {
		return tpl, err
	}
	var countries []shop.Country
	if err := p.Db.Select([]string{"id", "name"}).Order("name Asc").Find(&countries).Error; err != nil {
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
