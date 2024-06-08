package core

import (
	"sample/filters"
	"sample/utils"

	"github.com/billowdev/pagi"
	"github.com/gofiber/fiber/v2"
)

type (
	IHandlerInfs interface {
		HandleGetTodoes(c *fiber.Ctx) error
	}
	handlerDeps struct {
		todoSrv ISrvInfs
	}
)

func NewHandlers(
	todoSrv ISrvInfs,
) IHandlerInfs {
	return &handlerDeps{
		todoSrv: todoSrv,
	}
}

// HandleGetTodoes implements IHandlerInfs.
func (h *handlerDeps) HandleGetTodoes(c *fiber.Ctx) error {
	p := utils.InitPagination[filters.FTodo](c)
	p.Filters = filters.FTodo{
		Title:  c.Query("title"),
		Detail: c.Query("detail"),
		CommonTimeFilters: pagi.CommonTimeFilters{
			DateField:     c.Query("date_field"),
			CreatedAfter:  c.Query("created_after"),
			UpdatedAfter:  c.Query("updated_after"),
			CreatedBefore: c.Query("created_before"),
			UpdatedBefore: c.Query("updated_before"),
			CreatedAt:     c.Query("created_at"),
			UpdatedAt:     c.Query("updated_at"),
			StartDate:     c.Query("start_date"),
			EndDate:       c.Query("end_date"),
		},
	}
	result, err := h.todoSrv.GetTodoes(p)
	if err != nil {
		return utils.NewResponse[interface{}](c, fiber.StatusInternalServerError, "ERR_INTERNAL_SERVER", err.Error(), nil)
	}
	return utils.NewPaginationResponse(c, "success", result)
}
