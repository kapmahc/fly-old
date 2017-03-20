package erp

import (
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/engines/shop"
)

// Dao dao
type Dao struct {
	Db *gorm.DB `inject:""`
}

// ListCountry list country order by name asc
func (p *Dao) ListCountry() ([]shop.Country, error) {
	var items []shop.Country
	err := p.Db.Select([]string{"id", "name"}).Order("name ASC").Find(&items).Error
	return items, err
}

// ListZone list zone order by name asc
func (p *Dao) ListZone() ([]shop.Zone, error) {
	var items []shop.Zone
	err := p.Db.Select([]string{"id", "name", "active"}).Order("name ASC").Find(&items).Error
	return items, err
}

// ListState list state order by name asc
func (p *Dao) ListState() ([]shop.State, error) {
	var items []shop.State
	err := p.Db.Select([]string{"id", "name", "country_id", "zone_id"}).Order("name ASC").Find(&items).Error
	return items, err
}
