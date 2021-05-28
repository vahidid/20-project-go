package Models

import "gorm.io/gorm"

type Model interface {
	TableName() string
}

type Task struct {
	gorm.Model
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  bool   `json:"status"`
}

func (t *Task) TableName() string {
	return "tasks"
}
