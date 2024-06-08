package configs

import "github.com/billowdev/pagi"

type PagingResponse[T interface{}] struct {
	Code          int                `json:"code"`
	StatusCode    int                `json:"status_code"`
	StatusMessage string             `json:"status_message"`
	Data          interface{}        `json:"data"`
	Pagination    pagi.PagingInfo[T] `json:"pagination"`
}

