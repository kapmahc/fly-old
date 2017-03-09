package blog

import "time"

// Post post
type Post struct {
	Href      string    `json:"href"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Published time.Time `json:"published"`
}

// ByPublished sort
type ByPublished []Post

func (p ByPublished) Len() int {
	return len(p)
}

func (p ByPublished) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPublished) Less(i, j int) bool {
	return p[i].Published.UnixNano() > p[j].Published.UnixNano()
}
