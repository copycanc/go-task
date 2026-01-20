package internal

import (
	"go-br-task/internal/task"
	"go-br-task/internal/user"

	"github.com/gin-gonic/gin"
)

// Функция для инициализации эндпоинтов
func Init(r *gin.Engine, h *task.HandlerTask, u *user.HandlerUser) {
	// Получить все задачи
	r.GET("/tasks", h.GetTask)
	// Создать новую задачу
	r.POST("/tasks", h.CreateTask)
	// Получить задачу по ID
	r.GET("/tasks/:id", h.GetTaskID)
	// Удалить задачу по ID
	r.DELETE("/tasks/:id", h.DeleteTaskID)
	// Изменить статус задачи
	r.PUT("/tasks/:id", h.UpdateTaskID)

	// Получить всех пользователей
	r.GET("/users", u.GetUser)
	// Создать нового пользователя
	r.POST("/users", u.CreateUser)
	// Получить пользователя по ID
	r.GET("/users/:id", u.GetUserID)
	// Удалить пользователя по ID
	r.DELETE("/users/:id", u.DeleteUserID)
	// Изменить пользователя
	r.PUT("/users/:id", u.UpdateUserID)
}
