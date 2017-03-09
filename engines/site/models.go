package site

import (
	"time"

	"github.com/kapmahc/fly/web"
)

// Notice notice
type Notice struct {
	web.Model
	Body string
	Type string
}

// TableName table name
func (Notice) TableName() string {
	return "notices"
}

// LeaveWord leave-word
type LeaveWord struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"updatedAt"`
	Body      string
	Type      string
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}
