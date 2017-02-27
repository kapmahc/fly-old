package web

// Link link
type Link struct {
	Href  string
	Label string
}

// Dropdown dropdown
type Dropdown struct {
	Label string
	Href  string
	Items []Link
}
