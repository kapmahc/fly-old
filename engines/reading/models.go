package reading

import (
	"time"

	"github.com/kapmahc/fly/web"
)

// Book book
type Book struct {
	web.Model

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
