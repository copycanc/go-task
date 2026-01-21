package task

import (
	"go-br-task/internal/message"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerTask struct {
	taskService *TasksService
}

func NewHandler(taskService *TasksService) *HandlerTask {
	return &HandlerTask{taskService: taskService}
}

// Получить все задачи
func (h *HandlerTask) List(c *gin.Context) {
	taskList, httpStatus, err := h.taskService.GetAllTask()
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "OK",
		"tasks":  taskList,
	})
	return
}

// Создать новую задачу
func (h *HandlerTask) Create(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		message.StatusBadRequestDataH(c, err)
		return
	}
	if httpStatus, err := h.taskService.CreateTask(task); err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	message.StatusHttpSuccess(c)
	return
}

// Получить задачу по ID
func (h *HandlerTask) Get(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.taskService.TaskExist(id)
	if httpStatusE == 200 {
		task, httpStatus, err := h.taskService.GetTaskID(id)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "OK",
			"task":   task,
		})
		return
	}
	message.StatusHttpError(c, httpStatusE, errE)
	return
}

// Удалить задачу по ID
func (h *HandlerTask) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.DeleteTaskID(id)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		message.StatusHttpSuccess(c)
		return
	}
	message.StatusHttpError(c, httpStatus, err)
	return
}

// Изменить статус задачи
func (h *HandlerTask) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var taskUpdate TaskUpdate
	if errr := c.ShouldBindJSON(&taskUpdate); errr != nil {
		message.StatusBadRequestDataH(c, errr)
		return
	}
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.taskService.UpdateTaskID(id, taskUpdate.Status)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		message.StatusHttpSuccess(c)
		return
	}
	message.StatusHttpError(c, httpStatus, err)
	return
}
