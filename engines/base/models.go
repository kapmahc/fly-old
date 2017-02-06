package base

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Model base model
type Model struct {
	ID      uint `orm:"column(id)"`
	Updated time.Time
	Created time.Time
}

// Locale locale
type Locale struct {
	Model

	Code    string
	Message string
	Lang    string
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
}

func init() {
	orm.RegisterModel(&Locale{})
}
