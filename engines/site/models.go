package site

import (
	"time"

	"github.com/kapmahc/fly/engines/base"
)

// LeaveWord leave word
type LeaveWord struct {
	ID      uint
	Body    string
	Type    string
	Created time.Time
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}

// Notice notice
type Notice struct {
	base.Model

	Body string
	Type string
}

// TableName table name
func (Notice) TableName() string {
	return "notices"
}

// Link link
type Link struct {
	base.Model
	Loc       string
	Label     string
	Href      string
	SortOrder int
}

// TableName table name
func (Link) TableName() string {
	return "links"
}
