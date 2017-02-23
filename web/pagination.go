package web

// NewPagination new pagination
func NewPagination(href string, page, size, total int64) *Pagination {
	if size <= 0 || size >= 60 {
		size = 60
	}

	tp := total / size
	if total%size > 0 {
		tp = total/size + 1
	}
	if page <= 0 {
		page = 1
	}
	if page > tp {
		page = tp
	}

	var pages []int64
	for i := int64(1); i <= tp; i++ {
		if tp > 10 {
			if i < page-6 || i > page+6 {
				continue
			}
		}
		pages = append(pages, i)
	}

	return &Pagination{
		Href:       href,
		PageNo:     page,
		PageSize:   size,
		Pages:      pages,
		TotalCount: total,
		TotalPage:  tp,
		PrevPage:   page - 1,
		NextPage:   page + 1,
		Items:      make([]interface{}, 0),
	}
}

// Pagination pagination
type Pagination struct {
	Href       string
	PageNo     int64
	PageSize   int64
	TotalPage  int64
	TotalCount int64
	PrevPage   int64
	NextPage   int64
	Items      []interface{}
	Pages      []int64
}

// Limit limit
func (p *Pagination) Limit() int64 {
	return p.PageSize
}

// Offset offset
func (p *Pagination) Offset() int64 {
	return (p.PageNo - 1) * p.PageSize
}
