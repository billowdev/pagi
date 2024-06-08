package core

import (
	"sample/filters"
	"sample/models"

	"github.com/billowdev/pagi"
	"gorm.io/gorm"
)

type (
	IRepoInfs interface {
		GetTodoes(p pagi.PagingParams[filters.FTodo]) (*pagi.PagingInfo[[]models.TodoModel], error)
	}
	repoDeps struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) IRepoInfs {
	return &repoDeps{db: db}
}

// GetTodoes implements IRepoInfs.
func (r *repoDeps) GetTodoes(p pagi.PagingParams[filters.FTodo]) (*pagi.PagingInfo[[]models.TodoModel], error) {

	fp := p.Filters
	q := r.db
	q = pagi.ApplyFilter(q, "title", fp.Title, "contains")

	q = pagi.ApplyDatetimeFilters(q, pagi.CommonTimeFilters{
		CreatedAfter:  fp.CommonTimeFilters.CreatedAfter,
		UpdatedAfter:  fp.CommonTimeFilters.UpdatedAfter,
		CreatedBefore: fp.CommonTimeFilters.CreatedBefore,
		UpdatedBefore: fp.CommonTimeFilters.UpdatedBefore,
		CreatedAt:     fp.CommonTimeFilters.CreatedAt,
		UpdatedAt:     fp.CommonTimeFilters.UpdatedAt,
		StartDate:     fp.CommonTimeFilters.StartDate,
		EndDate:       fp.CommonTimeFilters.EndDate,
	})
	orderBy := pagi.NewOrderBy(pagi.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	q = q.Order(orderBy)
	pgR, err := pagi.Paginate[filters.FTodo, []models.TodoModel](p, q)
	if err != nil {
		return nil, err
	}
	return &pgR, nil
}
