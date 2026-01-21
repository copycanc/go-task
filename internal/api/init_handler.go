package api

import (
	"go-br-task/internal/auth"
	"go-br-task/internal/task"
	"go-br-task/internal/user"

	"github.com/gin-gonic/gin"
)

// Функция для инициализации эндпоинтов
func Init(r *gin.Engine, h *task.HandlerTask, u *user.HandlerUser, a *auth.HandlerAuth) {
	tasks := r.Group("/tasks")
	{
		tasks.GET("", h.List)
		tasks.POST("", h.Create)
		tasks.GET("/:id", h.Get)
		tasks.DELETE("/:id", h.Delete)
		tasks.PUT("/:id", h.Update)
	}

	users := r.Group("/users")
	{
		users.GET("", u.List)
		users.POST("", u.Create)
		users.GET("/:id", u.Get)
		users.DELETE("/:id", u.Delete)
		users.PUT("/:id", u.Update)
	}

	r.POST("/login")
	r.GET("/profile")

}
