package paging

type Page[T any] struct {
	Content       []T `json:"content"`
	*PageRequest  `json:"page_request"`
	TotalElements int  `json:"total_elements"`
	TotalPages    int  `json:"total_pages"`
	HasNext       bool `json:"has_next"`
	HasPrevious   bool `json:"has_previous"`
}

func NewPage[T any](content []T, pageable *PageRequest, totalElements int) *Page[T] {
	p := pageable

	hasPrevious := pageable.PageNumber > 1
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
