package orm

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Locale locale model
type Locale struct {
	ID        uint
	Lang      string
	Code      string
	Message   string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// I18nStore store for i18n
type I18nStore struct {
	Db *gorm.DB `inject:""`
}

// Get get
func (p *I18nStore) Get(lang, code string) (string, error) {
	var l Locale
	if err := p.Db.
		Select("message").
		Where("lang = ? AND code = ?", lang, code).
		First(&l).Error; err != nil {
		return "", err
	}
	return l.Message, nil
}

// Set set
func (p *I18nStore) Set(lang, code, message string) error {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lang, code).First(&l).RecordNotFound() {
		l.Lang = lang
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	return err
}

// Delete delete
func (p *I18nStore) Delete(lang, code string) error {
	return p.Db.Where("lang = ? AND code = ?", lang, code).Delete(Locale{}).Error
}

// All all items
func (p *I18nStore) All(lang string) (map[string]string, error) {
	var items []Locale
	err := p.Db.
		Select([]string{"code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").Find(&items).Error
	if err == nil {
		return nil, err
	}
	val := make(map[string]string)
	for _, l := range items {
		val[l.Code] = l.Message
	}
	return val, nil
}
