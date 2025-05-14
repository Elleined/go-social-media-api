package paging

import "errors"

type PageRequest struct {
	PageNumber int
	PageSize   int
}

func NewPageRequest(pageNumber int, pageSize int) (*PageRequest, error) {
	if pageNumber <= 0 {
		return nil, errors.New("page is required")
	}

	if pageSize <= 0 {
		return nil, errors.New("page size is required")
	}

	return &PageRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}

func (p PageRequest) Offset() int {
	return (p.PageNumber - 1) * p.PageSize
}
