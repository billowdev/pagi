# pagi: Pagination Helper Package

pagi is a Go package designed to simplify pagination in your applications. It provides utility functions and types to handle pagination parameters, sorting, filtering, and generating pagination links.

## Installation

To install pagi, use `go get`:

```bash 
go get -u github.com/billowdev/pagi
```

## Usage

### 1. Import the Package

```go
import "github.com/billowdev/pagi"
```

### 2. Pagination Parameters

The `PagingParams` type represents pagination parameters including limits, page numbers, sorting, and filtering.

### 3. Sorting

The `SortParams` type defines parameters for sorting, including the field to sort by, sorting order, and default sorting order.

### 4. Filtering

The `CommonTimeFilters` type contains common time-based filters such as date ranges and timestamps.

### 5. Utility Functions

#### AddWhereClauseIfNotEmpty

```go
func AddWhereClauseIfNotEmpty(query *gorm.DB, columnName string, filterValue string, filterType string) *gorm.DB
```

Adds a WHERE clause to the query if the filter value is not empty, based on the provided column name, filter value, and filter type.

#### ApplyFilter

```go
func ApplyFilter(query *gorm.DB, column string, value interface{}, filterType string) *gorm.DB
```

Applies filtering to the query based on the provided column, value, and filter type.

#### ApplyCommaFilter

```go
func ApplyCommaFilter(query *gorm.DB, columnName, filterValue string) *gorm.DB
```

Applies filtering to the query for comma-separated values in the specified column.

#### ApplyCommaFilterWithJoin

```go
func ApplyCommaFilterWithJoin(query *gorm.DB, joinTable, joinCondition, columnName, filterValue string) *gorm.DB
```

Applies filtering to the query with a JOIN condition for comma-separated values.

#### ApplyDatetimeFilters

```go
func ApplyDatetimeFilters(query *gorm.DB, filter CommonTimeFilters) *gorm.DB
```

Applies datetime filtering to the query based on the provided CommonTimeFilters.

#### ApplyDatetimePreloadFilters

```go
func ApplyDatetimePreloadFilters(query *gorm.DB, filter CommonTimeFilters, preloadKey string) *gorm.DB
```

Applies datetime filtering with preloading to the query based on the provided CommonTimeFilters.

## Examples

### Basic Pagination

```go
params := pagi.PagingParams{}
params.Limit = 10
params.Page = 1
params.Sort = "id"
params.Order = "asc"
params.BaseURL = "https://example.com"
```


## License

pagi is licensed under the MIT License. See the [LICENSE](https://github.com/billowdev/pagi/blob/main/LICENSE) file for details.



## Contributing

- Contributions are welcome! Feel free to open issues or submit pull requests on the [GitHub repository.](https://github.com/billowdev/pagi)
