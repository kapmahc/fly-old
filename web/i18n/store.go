package i18n

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
)

// Store i18n store
type Store interface {
	Set(l, k, v string) error
	Get(l, k string) (string, error)
	Del(lng, code string) error
	Codes(lang string) ([]string, error)
	Sync(dir string) error
}

// NewGormStore new gorm store
func NewGormStore(db *gorm.DB) Store {
	return &GormStore{db: db}
}

// GormStore gorm i18n store
type GormStore struct {
	db *gorm.DB
}

//Set set locale
func (p *GormStore) Set(lng string, code, message string) error {
	var l Locale
	var err error
	if p.db.Where("lang = ? AND code = ?", lng, code).First(&l).RecordNotFound() {
		l.Lang = lng
		l.Code = code
		l.Message = message
		err = p.db.Create(&l).Error
	} else {
		l.Message = message
		err = p.db.Save(&l).Error
	}
	return err
}

//Del del locale
func (p *GormStore) Del(lng, code string) error {
	return p.db.Where("lang = ? AND code = ?", lng, code).Delete(Locale{}).Error
}

// Get get
func (p *GormStore) Get(lng, code string) (string, error) {
	var l Locale
	if err := p.db.
		Select("message").
		Where("lang = ? AND code = ?", lng, code).
		First(&l).Error; err != nil {
		return "", err
	}

	return l.Message, nil
}

//Codes list locale keys
func (p *GormStore) Codes(lang string) ([]string, error) {
	var keys []string
	err := p.db.Model(&Locale{}).Where("lang = ?", lang).Pluck("code", &keys).Error
	return keys, err
}

// Sync sync records
func (p *GormStore) Sync(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		const ext = ".ini"
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ext {
			log.Debugf("Find locale file %s", path)
			if err != nil {
				return err
			}
			lang := name[0 : len(name)-len(ext)]
			if _, err := language.Parse(lang); err != nil {
				return err
			}
			cfg, err := ini.Load(path)
			if err != nil {
				return err
			}

			items := cfg.Section(ini.DEFAULT_SECTION).KeysHash()
			for k, v := range items {
				var l Locale
				if p.db.Where("lang = ? AND code = ?", lang, k).First(&l).RecordNotFound() {
					l.Lang = lang
					l.Code = k
					l.Message = v
					if err := p.db.Create(&l).Error; err != nil {
						return err
					}
				}
			}
			log.Infof("find %d items", len(items))
		}
		return nil
	})
}
