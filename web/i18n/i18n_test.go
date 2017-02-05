package i18n_test

import (
	"testing"

	"golang.org/x/text/language"

	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web/i18n"
	_ "github.com/mattn/go-sqlite3"
)

var lang = language.SimplifiedChinese.String()

func openDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil
}

func TestGormStore(t *testing.T) {
	db, err := openDatabase()
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&i18n.Locale{})
	db.Model(&i18n.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")

	s := i18n.NewGormStore(db)

	key := "hello"
	val := "你好"
	s.Set(lang, key, val)
	s.Set(lang, key+".1", val)

	ks, err := s.Codes(lang)
	if err != nil {
		t.Fatal(err)
	}
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	s.Del(lang, key)
}
