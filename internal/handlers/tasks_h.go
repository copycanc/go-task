package handlers

import (
	"go-br-task/internal/models"
	"go-br-task/utils/messages"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Получить все задачи
func (h *Handler) GetTask(c *gin.Context) {
	taskList, httpStatus, err := h.taskService.GetAllTask()
	if err != nil {
		messages.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "OK",
		"tasks":  taskList,
	})
	return
}

// Создать новую задачу
func (h *Handler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		messages.StatusBadRequestDataH(c, err)
		return
	}
	if httpStatus, err := h.taskService.CreateTask(task); err != nil {
		messages.StatusHttpError(c, httpStatus, err)
		return
	}
	messages.StatusHttpSuccess(c)
	return
}

// Получить задачу по ID
func (h *Handler) GetTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.taskService.TaskExist(id)
	if httpStatusE == 200 {
		task, httpStatus, err := h.taskService.GetTaskID(id)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "OK",
			"task":   task,
		})
		return
	}
	messages.StatusHttpError(c, httpStatusE, errE)
	return
}

// Удалить задачу по ID
func (h *Handler) DeleteTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.DeleteTaskID(id)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		messages.StatusHttpSuccess(c)
		return
	}
	messages.StatusHttpError(c, httpStatus, err)
	return
}

// Изменить статус задачи
func (h *Handler) UpdateTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var taskUpdate models.TaskUpdate
	if errr := c.ShouldBindJSON(&taskUpdate); errr != nil {
		messages.StatusBadRequestDataH(c, errr)
		return
	}
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.UpdateTaskID(id, taskUpdate.Status)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		messages.StatusHttpSuccess(c)
		return
	}
	messages.StatusHttpError(c, httpStatus, err)
	return
}
