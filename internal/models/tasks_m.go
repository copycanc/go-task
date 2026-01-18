package models

import (
	"time"

	"github.com/google/uuid"
)

// Структура тасок
type Task struct {
	ID          uuid.UUID
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      TaskStatus
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Структура для изменения статуса
type TaskUpdate struct {
	Status TaskStatus `json:"status" binding:"required"`
}

// Определяем константы для статусв
type TaskStatus string

const (
	NewT      TaskStatus = "Новая"
	Progress  TaskStatus = "В процессе"
	Completed TaskStatus = "Завершена"
)
