package web

import "time"

// H hash
type H map[string]interface{}

// K key type
type K string

const (
	// TypeMARKDOWN markdown format
	TypeMARKDOWN = "markdown"
	// TypeHTML html format
	TypeHTML = "html"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}

// Link link
type Link struct {
	Label string
	Href  string
}

// Dropdown dropdown
type Dropdown struct {
	Label string
	Items []Link
}
