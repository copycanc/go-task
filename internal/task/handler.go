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
}

// Получить задачу по ID
func (h *HandlerTask) Get(c *gin.Context) {
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	if !h.ensureTaskExist(c, id) {
		return
	}
	task, httpStatus, err := h.taskService.GetTaskID(id)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "OK",
		"task":   task,
	})
}

// Удалить задачу по ID
func (h *HandlerTask) Delete(c *gin.Context) {
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	if !h.ensureTaskExist(c, id) {
		return
	}
	httpStatus, err := h.taskService.DeleteTaskID(id)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	message.StatusHttpSuccess(c)
}

// Изменить статус задачи
func (h *HandlerTask) Update(c *gin.Context) {
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	var taskUpdate TaskUpdate
	if err := c.ShouldBindJSON(&taskUpdate); err != nil {
		message.StatusBadRequestDataH(c, err)
		return
	}
	if !h.ensureTaskExist(c, id) {
		return
	}
	httpStatus, err := h.taskService.UpdateTaskID(id, taskUpdate.Status)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	message.StatusHttpSuccess(c)
}

func parseUUIDParam(c *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(name))
	if err != nil {
		message.StatusBadRequestDataH(c, err)
		return uuid.Nil, false
	}
	return id, true
}

func (h *HandlerTask) ensureTaskExist(c *gin.Context, id uuid.UUID) bool {
	httpStatus, err := h.taskService.TaskExist(id)
	if httpStatus != 200 {
		message.StatusHttpError(c, httpStatus, err)
		return false
	}
	return true
}
