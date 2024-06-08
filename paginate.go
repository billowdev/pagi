package pagi

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

// Paginate is a function that facilitates pagination of data retrieved using GORM in Go.
// It takes in paging parameters and a GORM query, then returns a PagingInfo structure containing paginated data along with links for navigating through the pages.
//
// Parameters:
//   - p: PagingParams[FT] represents the paging parameters including page number, page size, base URL, etc.
//   - query: *gorm.DB is a pointer to the GORM query used to fetch data.
//
// Returns:
//   - PagingInfo[T]: A structure containing paginated data along with metadata and navigation links.
//   - error: An error, if any occurred during the pagination process.
//
// Usage Example:
//
//	// Define paging parameters
//	pagingParams := PagingParams{
//		Page:     1,
//		Limit:    10,
//		BaseURL:  "http://example.com/data",
//		Sort:     "created_at desc",
//	}
//
//	// Perform pagination
//	paginatedData, err := Paginate(pagingParams, db.Model(&YourModel{}))
//	if err != nil {
//		log.Fatal(err)
//	}
func Paginate[FT any, T any](p PagingParams[FT], query *gorm.DB) (PagingInfo[T], error) {
	var value T
	var totalRows int64
	if err := query.Model(value).Count(&totalRows).Error; err != nil {
		return PagingInfo[T]{}, err
	}
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	if err := query.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort()).Find(&value).Error; err != nil {
		return PagingInfo[T]{}, err

	}
	var nextLink, prevLink string
	if p.Page < p.TotalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page+1, p.Limit)
	}
	if p.Page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page-1, p.Limit)
	}
	return PagingInfo[T]{
		Links: PagingLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      p.TotalRows,
		Page:       p.GetPage(),
		PageSize:   p.GetLimit(),
		TotalPages: p.TotalPages,
		Rows:       value,
	}, nil
}

