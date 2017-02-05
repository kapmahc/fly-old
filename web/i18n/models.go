package i18n

import "time"

//Locale locale model
type Locale struct {
	ID        uint
	Lang      string
	Code      string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
}
