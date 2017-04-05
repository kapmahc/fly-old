package web

import "time"

const (
	// TypeMarkdown markdown format
	TypeMarkdown = "markdown"
	// TypeHTML html format
	TypeHTML = "html"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
