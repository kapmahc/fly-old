package web

const (
	// DATA data key
	DATA = "data"
	// NOTICE flash notice
	NOTICE = "notice"
	// ERROR flash error
	ERROR = "error"
	// ALERT flash alert
	ALERT = "alert"
)

// Link link
type Link struct {
	Model
	Href      string
	Label     string
	Loc       string
	SortOrder int
}

// Page page
type Page struct {
	Model
	Href      string
	Logo      string
	Title     string
	Summary   string
	Loc       string
	SortOrder int
}

// Dropdown dropdown
type Dropdown struct {
	Label     string
	Links     []Link
	SortOrder int
}
