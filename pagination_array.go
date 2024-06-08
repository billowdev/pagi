package pagi

import "fmt"

// PaginateArray is a function that paginates an array of elements of type T.
// It takes the array, page number, page size, and endpoint URL as input parameters, and returns a PagingInfo structure containing paginated data along with links for navigating through the pages, and the subset of data for the current page.
//
// Parameters:
//   - data: []T is the array of elements to be paginated.
//   - page: int represents the current page number.
//   - pageSize: int represents the number of elements per page.
//   - endpoint: string is the base URL endpoint used for constructing next and previous page links.
//
// Returns:
//   - PagingInfo[T]: A structure containing paginated data metadata and navigation links.
//   - []T: A subset of data representing the elements for the current page.
//
// Note:
//   - If pageSize or page is less than or equal to 0, the function returns an empty PagingInfo and nil slice.
//   - If the requested page number exceeds the total number of pages, the function returns an empty PagingInfo and nil slice.
//   - The function constructs next and previous page links based on the provided endpoint URL.
//   - The returned PagingInfo contains metadata such as total number of items, current page number, page size, and total number of pages.
func PaginateArray[T any](data []T, page, pageSize int, endpoint string) (PagingInfo[T], []T) {
	totalItems := len(data)
	if pageSize <= 0 || page <= 0 {
		return PagingInfo[T]{}, nil // or return an error
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	if page > totalPages {
		return PagingInfo[T]{}, nil // or return an error
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	pageData := data[start:end]

	// Construct next and previous links
	var nextLink, prevLink string
	if page < totalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page+1, pageSize)
	}
	if page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page-1, pageSize)
	}

	pagination := PagingInfo[T]{
		Links: PagingLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      int64(totalItems),
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
	return pagination, pageData
}
