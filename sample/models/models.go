package models

import (
	"time"

	"gorm.io/gorm"
)

type TodoModel struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null;size:255" json:"title"`
	Detail    string    `gorm:"not null;size:255" json:"detail"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var TNTodo = "todoes"

func (st *TodoModel) TableName() string {
	return TNTodo
}
