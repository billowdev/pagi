package pagi

import (
	"strconv"
)

// NewPaginationParams creates a new instance of PagingParams with the provided parameters.
//
// Parameters:
//   - host: The base URL for the pagination.
//   - sort: The field to sort the results by.
//   - order: The order in which to sort the results ('asc' or 'desc').
//   - page: The page number for pagination.
//   - pageSize: The number of items per page.
//   - defaultSort: The default field to sort the results by if 'sort' parameter is not provided.
//
// Returns:
//   - PagingParams[FilterType]: A PagingParams instance configured with the provided parameters.
//
// Example:
//
//	params := NewPaginationParams("https://example.com", "name", "asc", "1", "20", "id")
//	// Creates a PagingParams instance with base URL "https://example.com",
//	// sorting by "name" in ascending order, with 20 items per page and starting from page 1.
func NewPaginationParams[FilterType interface{}](host string, sort string, order string, page string, pageSize string, defaultSort string) PagingParams[FilterType] {
	defaultLimit := 10
	defaultPage := 1

	paginationParams := PagingParams[FilterType]{
		Limit:   defaultLimit,
		Page:    defaultPage,
		Sort:    defaultSort,
		BaseURL: host,
	}

	limit, err := strconv.Atoi(pageSize)
	if err == nil && limit > 0 {
		paginationParams.Limit = limit
	}

	pageInt, err := strconv.Atoi(page)
	if err == nil && pageInt > 0 {
		paginationParams.Page = pageInt
	}
	return paginationParams
}
