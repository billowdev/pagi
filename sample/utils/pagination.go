package utils

import (
	"fmt"
	"strings"

	"github.com/billowdev/pagi"
	"github.com/gofiber/fiber/v2"
)

type PaginationResponse struct {
	Links      pagi.PagingLinks `json:"links"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}

type APIPaginationResponse[T interface{}] struct {
	StatusCode    string             `json:"status_code"`
	StatusMessage string             `json:"status_message"`
	Data          T                  `json:"data"`
	Pagination    PaginationResponse `json:"pagination"`
}

func GetAPIEndpoint(c *fiber.Ctx) string {
	// Get the original URL from the request (excluding the query string)
	originalURL := c.OriginalURL()

	// Remove the query string (if any)
	if index := strings.Index(originalURL, "?"); index != -1 {
		originalURL = originalURL[:index]
	}

	// Construct the desired URL with the current host and path
	return fmt.Sprintf("https://%s%s", c.Hostname(), originalURL)
}

func InitPagination[FilterType interface{}](c *fiber.Ctx) pagi.PagingParams[FilterType] {
	host := GetAPIEndpoint(c) // Assuming utils.GetAPIEndpoint is a function you have
	defaultSort := "created_at desc"
	sort := c.Query("sort")
	order := c.Query("order")
	page := c.Query("page")
	limit := c.Query("limit")
	return pagi.InitPagingParams[FilterType](host, sort, order, page, limit, defaultSort)
}

func GetPaginationInfo[T interface{}](payload pagi.PagingInfo[T]) PaginationResponse {
	return PaginationResponse{
		Links:      payload.Links,
		Total:      payload.Total,
		Page:       payload.Page,
		PageSize:   payload.PageSize,
		TotalPages: payload.TotalPages,
	}
}

func NewPaginationResponse[T interface{}](c *fiber.Ctx, message string, data pagi.PagingInfo[T]) error {
	response := APIPaginationResponse[T]{
		StatusCode:    "200",
		StatusMessage: message,
		Data:          data.Rows,
		Pagination:    GetPaginationInfo[T](data),
	}
	return c.Status(200).JSON(response)
}
