package web

// Link link
type Link struct {
	Model
	Href      string
	Label     string
	Loc       string
	SortOrder int
}

// TableName table name
func (Link) TableName() string {
	return "links"
}

// Page page
type Page struct {
	Model
	Href      string
	Logo      string
	Title     string
	Summary   string
	Action    string
	Loc       string
	SortOrder int
}

// TableName table name
func (Page) TableName() string {
	return "pages"
}

// Dropdown dropdown
type Dropdown struct {
	Label     string
	Links     []*Link
	SortOrder int
}
