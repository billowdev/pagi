package pagi

import (
	"fmt"
	"strings"
)

// NewOrderBy constructs the ORDER BY clause based on sorting parameters.
//
// Parameters:
//   - params: A SortParams instance containing sorting parameters.
//
// Returns:
//   - string: The constructed ORDER BY clause.
//
// Example:
//
//	orderBy := NewOrderBy(SortParams{"name", "asc", "id asc"})
//	// Constructs an ORDER BY clause sorting by "name" in ascending order,
//	// and if invalid sorting parameters are provided, falls back to default order.
func NewOrderBy(params SortParams) string {
	// DefaultOrderBy represents the default ORDER BY clause
	// Validate and sanitize sorting parameters
	sortField := strings.TrimSpace(params.Sort)
	sortOrder := strings.ToUpper(strings.TrimSpace(params.Order))

	// Construct the ORDER BY clause based on parameters
	if sortField != "" {
		if sortOrder == "ASC" || sortOrder == "DESC" {
			return fmt.Sprintf("%s %s", sortField, sortOrder)
		}
	}

	// If sorting parameters are invalid or not provided, fallback to default order
	return params.DefaultOrderBy
}

// GetLimit returns the limit value for pagination. If no limit is set, it defaults to 10.
//
// Returns:
//   - int: The limit value for pagination.
func (p *PagingParams[FilterType]) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

// GetPage returns the current page number for pagination. If no page is set, it defaults to 1.
//
// Returns:
//   - int: The current page number for pagination.
func (p *PagingParams[FilterType]) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

// GetOffset calculates and returns the offset for pagination based on the current page and limit.
//
// Returns:
//   - int: The offset value for pagination.
func (p *PagingParams[FilterType]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetSort returns the sort field and order for pagination. If no sorting parameters are provided, it defaults to "Id desc".
//
// Returns:
//   - string: The sort field and order for pagination.
func (p *PagingParams[FilterType]) GetSort() string {
	if p.Sort == "" {
		return "Id desc"
	}
	return p.Sort
}

// GetAPIEndpoint constructs the API endpoint URL based on the provided host and path.
//
// Parameters:
//   - host: The host of the API.
//   - path: The path of the API endpoint.
//
// Returns:
//   - string: The constructed API endpoint URL.
//
// Example:
//
//	endpoint := GetAPIEndpoint("example.com", "/api/v1/resource")
//	// Constructs the API endpoint URL as "https://example.com/api/v1/resource".
func GetAPIEndpoint(host, path string) string {
	// Remove the query string (if any)
	if index := strings.Index(path, "?"); index != -1 {
		path = path[:index]
	}
	// Construct the desired URL with the given host and path
	return fmt.Sprintf("https://%s%s", host, path)
}
