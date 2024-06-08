package utils

import "github.com/gofiber/fiber/v2"

type APIResponse[T interface{}] struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Data          T      `json:"data"`
}

func NewResponse[T interface{}](c *fiber.Ctx, http_status int, status_code string, message string, data T) error {
	response := APIResponse[T]{
		StatusCode:    status_code,
		StatusMessage: message,
		Data:          data,
	}
	return c.Status(http_status).JSON(response)
}
