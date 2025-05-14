package paging

type Page[T any] struct {
	Content       []T `json:"content"`
	*PageRequest  `json:"page_request"`
	TotalElements int  `json:"total_elements"`
	TotalPages    int  `json:"total_pages"`
	HasNext       bool `json:"has_next"`
	HasPrevious   bool `json:"has_previous"`
}

func NewPage[T any](content []T, pageRequest *PageRequest, totalElements int) *Page[T] {
	p := pageRequest

	hasPrevious := pageRequest.PageNumber > 1
	hasNext := p.Offset()+p.PageSize < totalElements
	totalPages := (totalElements + p.PageSize - 1) / p.PageSize

	return &Page[T]{
		Content:       content,
		PageRequest:   p,
		TotalElements: totalElements,
		TotalPages:    totalPages,
		HasNext:       hasNext,
		HasPrevious:   hasPrevious,
	}
}
