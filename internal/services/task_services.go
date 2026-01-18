package services

import (
	"errors"
	"go-br-task/internal/interfaces"
	"go-br-task/internal/models"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TasksService struct {
	storage interfaces.TasksStorage
}

func NewTaskService(storage interfaces.TasksStorage) *TasksService {
	return &TasksService{
		storage: storage,
	}
}

func (s *TasksService) GetAllTask() (map[uuid.UUID]models.Task, error) {
	task, err := s.storage.GetAllTask()
	if err != nil {
		slog.Error("Ошибка", err)
		return nil, errors.New("ошибка при получении данных")
	}
	return task, nil
}

func (s *TasksService) CreateTask(task models.Task) error {
	task = models.Task{
		ID:          uuid.New(),
		Title:       task.Title,
		Description: task.Description,
		Status:      models.NewT,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	if err := s.storage.SaveTask(task); err != nil {
		slog.Error("Ошибка", err)
		return errors.New("ошибка при сохранении")
	}
	return nil
}

func (s *TasksService) GetTaskID(uuid uuid.UUID) (*models.Task, int, error) {
	task, err := s.storage.GetTaskID(uuid)
	if err != nil {
		slog.Error("Ошибка", err)
		return nil, http.StatusInternalServerError, errors.New("ошибка при получении данных")
	}
	return task, http.StatusOK, nil
}

func (s *TasksService) TaskExist(uuid uuid.UUID) (int, error) {
	exist, err := s.storage.ExistTask(uuid)
	if !exist {
		return http.StatusNotFound, errors.New("задача не найдена")
	}
	if err != nil {
		slog.Error("Ошибка", err)
		return http.StatusInternalServerError, errors.New("ошибка при получении данных")
	}
	return 200, nil
}

func (s *TasksService) DeleteTaskID(uuid uuid.UUID) (int, error) {
	if err := s.storage.DeleteTask(uuid); err != nil {
		slog.Error("Ошибка", err)
		return http.StatusInternalServerError, errors.New("ошибка при удалении данных")
	}
	return http.StatusOK, nil
}

func (s *TasksService) UpdateTaskID(uuid uuid.UUID, status models.TaskStatus) (int, error) {
	task, err := s.storage.GetTaskID(uuid)
	if err != nil {
		slog.Error("Ошибка", err)
		return http.StatusInternalServerError, errors.New("ошибка при получении данных")
	}
	switch status {
	case models.Progress, models.NewT:
		task.Status = status
		if err = s.storage.SaveTask(*task); err != nil {
			slog.Error("Ошибка", err)
			return http.StatusInternalServerError, errors.New("ошибка при обновлении статуса")
		}
		return http.StatusOK, nil
	case models.Completed:
		task.Status = status
		now := time.Now()
		task.CompletedAt = &now
		if err = s.storage.SaveTask(*task); err != nil {
			slog.Error("Ошибка", err)
			return http.StatusInternalServerError, errors.New("ошибка при обновлении статуса")
		}
		return http.StatusOK, nil
	default:
		return http.StatusBadRequest, errors.New("данного статуса не существует")
	}
}
