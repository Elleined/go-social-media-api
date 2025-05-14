package paging

import (
	"errors"
	"strconv"
)

type PageRequest struct {
	PageNumber int
	PageSize   int
}

func NewPageRequestStr(pageNumber, pageSize string) (*PageRequest, error) {
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, err
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}

	pageRequest, err := NewPageRequest(pageNumberInt, pageSizeInt)
	if err != nil {
		return nil, err
	}

	return pageRequest, nil
}

func NewPageRequest(pageNumber, pageSize int) (*PageRequest, error) {
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
