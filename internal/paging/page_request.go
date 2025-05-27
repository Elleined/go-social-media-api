package paging

import (
	"errors"
	"strconv"
	"strings"
)

type PageRequest struct {
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
	Field      string `json:"field"`
	SortBy     string `json:"sort_by"`
}

func NewPageRequestStr(pageNumber, pageSize, field, sortBy string) (*PageRequest, error) {
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, err
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}

	pageRequest, err := NewPageRequest(pageNumberInt, pageSizeInt, field, sortBy)
	if err != nil {
		return nil, err
	}

	return pageRequest, nil
}

func NewPageRequest(pageNumber, pageSize int, field, sortBy string) (*PageRequest, error) {
	if pageNumber <= 0 {
		return nil, errors.New("page is required")
	}

	if pageSize <= 0 {
		return nil, errors.New("page size is required")
	}

	trimmedField := strings.TrimSpace(field)
	if trimmedField == "" {
		return nil, errors.New("field cannot be empty")
	}

	trimmedSortBy := strings.TrimSpace(sortBy)
	if trimmedSortBy == "" {
		return nil, errors.New("sortBy cannot be empty")
	}

	return &PageRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Field:      strings.ToLower(trimmedField),
		SortBy:     strings.ToUpper(trimmedSortBy),
	}, nil
}

func (p PageRequest) Offset() int {
	return (p.PageNumber - 1) * p.PageSize
}
