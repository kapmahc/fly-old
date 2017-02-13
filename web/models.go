package web

// Link link
type Link struct {
	Label string
	Href  string
}

// Dropdown dropdown
type Dropdown struct {
	Label string
	Items []Link
}
