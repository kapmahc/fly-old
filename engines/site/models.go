package site

import (
	"time"

	"github.com/kapmahc/sky"
)

// Post post
type Post struct {
	sky.Model
	Key   string
	Lang  string
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
	sky.Model

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
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}
