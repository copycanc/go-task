package task

import (
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type TasksService struct {
	storage TasksStorage
}

func NewTaskService(storage TasksStorage) *TasksService {
	return &TasksService{
		storage: storage,
	}
}

func (s *TasksService) GetAllTask() (map[uuid.UUID]Task, int, error) {
	task, err := s.storage.GetAllTask()
	if err != nil {
		slog.Error("Ошибка", err)
		return nil, 500, errors.New("ошибка при получении данных")
	}
	return task, 200, nil
}

func (s *TasksService) CreateTask(task Task) (int, error) {
	task = Task{
		ID:          uuid.New(),
		Title:       task.Title,
		Description: task.Description,
		Status:      NewT,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	if err := s.storage.SaveTask(task); err != nil {
		slog.Error("Ошибка", err)
		return 500, errors.New("ошибка при сохранении")
	}
	return 200, nil
}

func (s *TasksService) GetTaskID(uuid uuid.UUID) (*Task, int, error) {
	task, err := s.storage.GetTaskID(uuid)
	if err != nil {
		slog.Error("Ошибка", err)
		return nil, 500, errors.New("ошибка при получении данных")
	}
	return task, 200, nil
}

func (s *TasksService) TaskExist(uuid uuid.UUID) (int, error) {
	exist, err := s.storage.ExistTask(uuid)
	if !exist {
		return 404, errors.New("задача не найдена")
	}
	if err != nil {
		slog.Error("Ошибка", err)
		return 500, errors.New("ошибка при получении данных")
	}
	return 200, nil
}

func (s *TasksService) DeleteTaskID(uuid uuid.UUID) (int, error) {
	if err := s.storage.DeleteTask(uuid); err != nil {
		slog.Error("Ошибка", err)
		return 500, errors.New("ошибка при удалении данных")
	}
	return 200, nil
}

func (s *TasksService) UpdateTaskID(uuid uuid.UUID, status TaskStatus) (int, error) {
	task, err := s.storage.GetTaskID(uuid)
	if err != nil {
		slog.Error("Ошибка", err)
		return 500, errors.New("ошибка при получении данных")
	}
	switch status {
	case Progress, NewT:
		task.Status = status
		if err = s.storage.SaveTask(*task); err != nil {
			slog.Error("Ошибка", err)
			return 500, errors.New("ошибка при обновлении статуса")
		}
		return 200, nil
	case Completed:
		task.Status = status
		now := time.Now()
		task.CompletedAt = &now
		if err = s.storage.SaveTask(*task); err != nil {
			slog.Error("Ошибка", err)
			return 500, errors.New("ошибка при обновлении статуса")
		}
		return 200, nil
	default:
		return 400, errors.New("данного статуса не существует")
	}
}
