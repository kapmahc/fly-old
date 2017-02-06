package base

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Locale locale
type Locale struct {
	ID        uint `orm:"column(id)"`
	Code      string
	Message   string
	Lang      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
}

func init() {
	orm.RegisterModel(&Locale{})
}
