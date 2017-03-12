package web

import (
	"net/http"
	"strconv"
)

// NewPagination new pagination
func NewPagination(r *http.Request, total int64) *Pagination {
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)

	if size <= 0 || size >= 60 {
		size = 60
	}
	if page <= 0 {
		page = 1
	}
	if page*size > total {
		page = total / size
		if total%size != 0 {
			page++
		}
	}

	var count = total / size
	if total == 0 {
		count = 1
	}
	if total%page != 0 {
		count++
	}

	return &Pagination{
		Page:  page,
		Size:  size,
		Total: total,
		Count: count,
		Items: make([]interface{}, 0),
	}
}

// Pagination pagination
type Pagination struct {
	Page  int64         `json:"page"`
	Size  int64         `json:"size"`
	Total int64         `json:"total"`
	Count int64         `json:"count"`
	Items []interface{} `json:"items"`
}

// Limit limit
func (p *Pagination) Limit() int64 {
	return p.Size
}

// Offset offset
func (p *Pagination) Offset() int64 {
	return (p.Page - 1) * p.Size
}
