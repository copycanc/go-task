package task

import "github.com/google/uuid"

type TasksStorage interface {
	GetAllTask() (map[uuid.UUID]Task, error)
	SaveTask(Task) error
	GetTaskID(id uuid.UUID) (*Task, error)
	ExistTask(id uuid.UUID) (bool, error)
	DeleteTask(id uuid.UUID) error
}
