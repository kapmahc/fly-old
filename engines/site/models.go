package site

import (
	"time"

	"github.com/kapmahc/sky"
)

// Model model
type Model struct {
	sky.Model

	SiteID uint
	Site   Site
}

// Site site
type Site struct {
	sky.Model

	Title       string
	SubTitle    string
	Description string
	Keywords    string
	Copyright   string
}

// Host host
type Host struct {
	Model

	Name string
}

// TableName table name
func (Host) TableName() string {
	return "hosts"
}

// Post post
type Post struct {
	Model

	Title string
	Body  string
	Type  string
}

// TableName table name
func (Post) TableName() string {
	return "posts"
}

// Notice notice
type Notice struct {
	Model

	Body string `json:"body"`
	Type string `json:"type"`
}

// TableName table name
func (Notice) TableName() string {
	return "notices"
}

// LeaveWord leave-word
type LeaveWord struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`

	SiteID uint
	Site   Site
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}
