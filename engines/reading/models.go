package reading

import (
	"time"

	"github.com/kapmahc/fly/web"
)

// Book book
type Book struct {
	web.Model

	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Lang        string    `json:"lang"`
	File        string    `json:"-"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Cover       string    `json:"cover"`
}

// TableName table name
func (Book) TableName() string {
	return "reading_books"
}