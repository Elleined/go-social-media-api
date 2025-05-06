package utils

import (
	"strconv"
)

// PageSize == limit
// Offset == offset
func Paginate(page, pageSize string) (limit, offset int, err error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	limit, err = strconv.Atoi(pageSize)
	if err != nil {
		return 0, 0, err
	}

	offset = (pageInt - 1) * limit
	return limit, offset, nil
}
