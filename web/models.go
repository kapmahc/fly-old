package web

import "time"

const (
	// TypeMARKDOWN markdown format
	TypeMARKDOWN = "markdown"
	// TypeHTML html format
	TypeHTML = "html"
	// DATA data key
	DATA = "data"
)

//Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"updatedAt"`
	UpdatedAt time.Time `json:"createdAt"`
}
