package handlers

import (
	"go-br-task/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskService *services.TasksService
}
type HandlerUser struct {
	userService *services.UserService
}

func NewHandlerUser(userService *services.UserService) *HandlerUser {
	return &HandlerUser{userService: userService}
}
func NewHandler(taskService *services.TasksService) *Handler {
	return &Handler{taskService: taskService}
}

// Функция для инициализации эндпоинтов
func Init(r *gin.Engine, h *Handler, u *HandlerUser) {
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
