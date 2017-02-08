package reading

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/base"
)

// Book book
type Book struct {
	base.Model

	Author      string
	Publisher   string
	Title       string
	Type        string
	Lang        string
	File        string
	Subject     string
	Description string
	PublishedAt time.Time
	Cover       string
}

// TableName table name
func (Book) TableName() string {
	return "reading_books"
}

func init() {
	orm.RegisterModel(&Book{})
}
