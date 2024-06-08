package pagi

// PagingParams represents pagination parameters including limits, sorting, total rows, and filters.
type PagingParams[FilterType interface{}] struct {
	Limit      int        `json:"limit"`
	Page       int        `json:"page"`
	Sort       string     `json:"sort"`
	Order      string     `json:"order"`
	TotalRows  int64      `json:"total_rows"`
	TotalPages int        `json:"total_pages"`
	BaseURL    string     `json:"base_url"`
	Filters    FilterType `json:"filters"`
}

// SortParams contains parameters for sorting.
type SortParams struct {
	Sort           string // Field to sort by
	Order          string // Sorting order ("ASC" or "DESC")
	DefaultOrderBy string // Default
}

// PagingLinks contains links for next and previous pages.
type PagingLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// PagingInfo provides information about pagination, including links, total count, page number, page size, total pages, and rows.
type PagingInfo[T interface{}] struct {
	Links      PagingLinks `json:"links"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	Rows       T           `json:"rows,omitempty"`
}

// CommonTimeFilters contains common time-based filters such as date ranges and timestamps.
type CommonTimeFilters struct {
	DateField     string `json:"date_field"`
	CreatedAfter  string `json:"created_after"`
	UpdatedAfter  string `json:"updated_after"`
	CreatedBefore string `json:"created_before"`
	UpdatedBefore string `json:"updated_before"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}
