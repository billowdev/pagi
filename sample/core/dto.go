package core

import "time"

type TodoDto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
