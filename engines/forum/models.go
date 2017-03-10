package forum

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

// Article article
type Article struct {
	web.Model
	Title   string
	Summary string
	Body    string
	Type    string

	UserID uint
	User   auth.User
	Tags   []Tag `gorm:"many2many:forum_articles_tags;"`
}

// TableName table name
func (Article) TableName() string {
	return "forum_articles"
}

// Tag tag
type Tag struct {
	web.Model
	Name     string
	Articles []Article `gorm:"many2many:forum_articles_tags;"`
}

// TableName table name
func (Tag) TableName() string {
	return "forum_tags"
}

// Comment comment
type Comment struct {
	web.Model
	Body string
	Type string

	UserID    uint
	User      auth.User
	ArticleID uint
	Article   Article
}

// TableName table name
func (Comment) TableName() string {
	return "forum_comments"
}
