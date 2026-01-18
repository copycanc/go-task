package handlers

import (
	"go-br-task/internal/models"
	"go-br-task/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Получить все задачи
func (h *Handler) GetTask(c *gin.Context) {
	taskList, err := h.taskService.GetAllTask()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"status": "ok",
		"tasks":  taskList,
	})
	return
}

// Создать новую задачу
func (h *Handler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		messages.IncorrectData(c, err)
		return
	}
	if err := h.taskService.CreateTask(task); err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"status": "ok",
		"text":   "Задача создана",
	})
	return
}

// Получить задачу по ID
func (h *Handler) GetTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.taskService.TaskExist(id)
	if httpStatusE == 200 {
		task, httpStatus, err := h.taskService.GetTaskID(id)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "ok",
			"task":   task,
		})
		return
	}
	c.JSONP(httpStatusE, gin.H{
		"status":  "error",
		"message": errE.Error(),
	})
	return
}

// Удалить задачу по ID
func (h *Handler) DeleteTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.DeleteTaskID(id)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status":  "ok",
			"message": "Задача удалена",
		})
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	return
}

// Изменить статус задачи
func (h *Handler) UpdateTaskID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var taskUpdate models.TaskUpdate
	if errr := c.ShouldBindJSON(&taskUpdate); errr != nil {
		messages.IncorrectData(c, errr)
		return
	}
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.UpdateTaskID(id, taskUpdate.Status)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status":  "ok",
			"message": "Задача обновлена",
		})
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	return
}
