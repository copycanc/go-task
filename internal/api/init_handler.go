package api

import (
	"go-br-task/internal/task"
	"go-br-task/internal/user"

	"github.com/gin-gonic/gin"
)

// Функция для инициализации эндпоинтов
func Init(r *gin.Engine, h *task.HandlerTask, u *user.HandlerUser) {
	tasks := r.Group("/task")
	{
		// Получить все задачи
		tasks.GET("", h.List)
		// Создать новую задачу
		tasks.POST("", h.Create)
		// Получить задачу по ID
		tasks.GET("/:id", h.Get)
		// Удалить задачу по ID
		tasks.DELETE("/:id", h.Delete)
		// Изменить статус задачи
		tasks.PUT("/:id", h.Update)
	}

	users := r.Group("/user")
	{
		// Получить всех пользователей
		users.GET("", u.List)
		// Создать нового пользователя
		users.POST("", u.Create)
		// Получить пользователя по ID
		users.GET("/:id", u.Get)
		// Удалить пользователя по ID
		users.DELETE("/:id", u.Delete)
		// Изменить пользователя
		users.PUT("/:id", u.Update)
	}

}
