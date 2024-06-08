package core

import (
	"errors"
	"sample/filters"

	"github.com/billowdev/pagi"
)

type (
	ISrvInfs interface {
		GetTodoes(p pagi.PagingParams[filters.FTodo]) (pagi.PagingInfo[[]TodoDto], error)
	}
	srvDeps struct {
		todoRepo IRepoInfs
	}
)

func NewServices(
	todoRepo IRepoInfs,
) ISrvInfs {
	return &srvDeps{
		todoRepo: todoRepo,
	}
}

// GetTodoes implements ISrvInfs.
func (s *srvDeps) GetTodoes(p pagi.PagingParams[filters.FTodo]) (pagi.PagingInfo[[]TodoDto], error) {
	pgR, err := s.todoRepo.GetTodoes(p)
	if err != nil {
		return pagi.PagingInfo[[]TodoDto]{}, err
	}
	if pgR == nil {
		err := errors.New("no data found")
		return pagi.PagingInfo[[]TodoDto]{}, err
	}

	var resp []TodoDto
	for _, v := range pgR.Rows {
		resp = append(resp, TodoDto{
			ID:        v.ID,
			Title:     v.Title,
			Detail:    v.Detail,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return pagi.PagingInfo[[]TodoDto]{
		Links:      pgR.Links,
		Total:      pgR.Total,
		Page:       pgR.Page,
		PageSize:   pgR.PageSize,
		TotalPages: pgR.TotalPages,
		Rows:       resp,
	}, nil
}
