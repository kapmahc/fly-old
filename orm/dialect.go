package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/migrate"
)

// MigrateDialect migrate dialect
type MigrateDialect struct {
	Db *gorm.DB `inject:""`
}

// Up migrate
func (p *MigrateDialect) Up(file string) error {
	return p.exec(file, true)
}

// Down rollback
func (p *MigrateDialect) Down(file string) error {
	return p.exec(file, false)
}

// Check check tables
func (p *MigrateDialect) Check() {
	var m migrate.Model
	p.Db.AutoMigrate(&m)
}

func (p *MigrateDialect) exec(file string, up bool) error {
	m := migrate.Model{File: file, Applied: up}
	var sql string
	var err error
	if m.Applied {
		sql, err = m.Up()
	} else {
		sql, err = m.Down()
	}
	if err != nil {
		return err
	}
	if err := p.Db.Exec(sql).Error; err != nil {
		return err
	}
	if p.Db.Where("file = ?", m.File).First(&m).RecordNotFound() {
		return p.Db.Create(&m).Error
	}
	return p.Db.Where("file = ?", m.File).Update("apply", m.Applied).Error

}

// Version current version
func (p *MigrateDialect) Version() (*migrate.Model, error) {
	var m migrate.Model
	if err := p.Db.Where("apply", true).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// Exec exec sql
func (p *MigrateDialect) Exec(sql string) error {
	return p.Db.Exec(sql).Error
}

// All list all migrations
func (p *MigrateDialect) All() ([]migrate.Model, error) {
	var items []migrate.Model
	err := p.Db.Order("id ASC").Find(&items).Error
	return items, err
}
